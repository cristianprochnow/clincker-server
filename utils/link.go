package utils

type LinkUtil struct {
	GenerateHash func(input string) string
}

func Link() LinkUtil {
	return LinkUtil{
		GenerateHash: generateLinkHash,
	}
}

func generateLinkHash(input string) string {
	return Crypto().Serial(input, 32)
}
