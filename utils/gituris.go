package utils

func GitURIs(dirPath string) (gitRepoURIs []string, err error) {
	var gru string
	gitRepoURIs = []string{}
	for _, grd := range GitDirs(dirPath) {
		if gru, err = GitURI(grd); err != nil {
			return
		}
		gitRepoURIs = append(gitRepoURIs, gru)
	}
	return
}
