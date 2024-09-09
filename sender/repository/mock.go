package repository

type SenderRepoMock struct {
}

var DefaultMock = 	SenderRepoMock{}

// GetUserId implements SendRepo.
func (s *SenderRepoMock) GetUserId(mail string) (int, error) {
	return 100, nil
}
