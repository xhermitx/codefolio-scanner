package servers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
	"github.com/xhermitx/gitpulse-tracker/github-service/API"
	"github.com/xhermitx/gitpulse-tracker/github-service/models"
)

func FetchData(w http.ResponseWriter, r *http.Request) {

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error processing the request", http.StatusBadRequest)
	}
	defer r.Body.Close()

	var res models.Job

	if err = json.Unmarshal(reqBody, &res); err != nil {
		http.Error(w, "Error processing the request", http.StatusBadRequest)
	}

	if len(res.Usernames) == 0 {
		http.Error(w, "Error processing the request", http.StatusBadRequest)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(res.Usernames))

	var candidate models.Candidate
	// GET EACH CANDIDATE'S DATA FROM GITHUB
	for _, u := range res.Usernames {
		profile, err := API.GetUserDetails(u)
		if err != nil {
			log.Println(err)
		} else {
			candidate = models.Candidate{
				JobId:         res.JobID,
				GithubId:      profile.Data.User.Login,
				Followers:     uint(profile.Data.User.Followers.TotalCount),
				Contributions: uint(profile.Data.User.ContributionsCollection.ContributionCalendar.TotalContributions),
				MostPopularRepo: func() string {
					if len(profile.Data.User.Repositories.Nodes) > 0 {
						return profile.Data.User.Repositories.Nodes[0].Name
					}
					return ""
				}(),
				RepoStars: func() uint {
					if len(profile.Data.User.Repositories.Nodes) > 0 {
						return uint(profile.Data.User.Repositories.Nodes[0].StargazerCount)
					}
					return 0
				}(),
			}

			// candidate := models.Candidate{
			// 	JobId:           2,
			// 	GithubId:        u,
			// 	Followers:       uint(4 + i), // Added a variable for different scores on redis
			// 	Contributions:   20,
			// 	MostPopularRepo: "Test",
			// 	RepoStars:       200,
			// 	Status:          false,
			// }

			// CREATE A GO ROUTINE FOR EACH PUBLISH ON THE QUEUE
			go func(candidate models.Candidate) {
				defer wg.Done()
				if err = API.Publish(candidate); err != nil {
					fmt.Print(err)
				}
			}(candidate)

		}

		wg.Wait()
		if err = API.Publish(models.Candidate{JobId: candidate.JobId, Status: true}); err != nil {
			log.Print(err)
		}
	}
}

func HttpServer() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/github", FetchData).Methods("POST")

	log.Fatal(http.ListenAndServe(os.Getenv("GITHUB_ADDRESS"), router))
}
