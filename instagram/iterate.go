package instagram

// IterateMedia makes pagination easy by converting the repeated api.NextMedias() call to a channel of media.
// If you desire to break early, pass in a doneChan and close when you are breaking from iteration
func (api *Api) IterateMedia(res *PaginatedMediasResponse, doneChan <-chan bool) (<-chan *Media, <-chan error) {
	mediaChan := make(chan *Media)
	errChan := make(chan error, 1)

	go func() {
		defer close(mediaChan)
		defer close(errChan)

		for {
			if res == nil {
				return
			}

			if len(res.Medias) == 0 {
				// No more Medias
				return
			}

			// Iterate
			for i := 0; i < len(res.Medias); i++ {
				select {
				case mediaChan <- &res.Medias[i]:
				case <-doneChan:
					return
				}
			}

			// Paginate to next response
			var err error
			res, err = api.NextMedias(res.Pagination)
			if err != nil {
				errChan <- err
				return
			}
		}
	}()

	return mediaChan, errChan
}

// IterateUsers makes pagination easy by converting the repeated api.NextUsers() call to a channel of users.
// If you desire to break early, pass in a doneChan and close when you are breaking from iteration
func (api *Api) IterateUsers(res *PaginatedUsersResponse, doneChan <-chan bool) (<-chan *User, <-chan error) {
	userChan := make(chan *User)
	errChan := make(chan error)

	go func() {
		defer close(userChan)
		defer close(errChan)

		for {
			if res == nil {
				return
			}

			if len(res.Users) == 0 {
				// No more users
				return
			}

			// Iterate
			for i := 0; i < len(res.Users); i++ {
				select {
				case userChan <- &res.Users[i]:
				case <-doneChan:
					return
				}
			}

			// Paginate to next response
			var err error
			res, err = api.NextUsers(res.Pagination)
			if err != nil {
				errChan <- err
				return
			}
		}
	}()

	return userChan, errChan
}
