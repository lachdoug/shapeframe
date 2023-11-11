package utils

func GitRepoURIs(dirPath string) (gitRepoURIs []string, err error) {
	var gru string
	gitRepoURIs = []string{}
	for _, grd := range GitRepoDirs(dirPath) {
		if gru, err = GitRepoURI(grd); err != nil {
			return
		}
		gitRepoURIs = append(gitRepoURIs, gru)
	}
	return
}
