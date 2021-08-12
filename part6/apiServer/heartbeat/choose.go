package heartbeat

import "math/rand"

func ChooseRandomDataServers(n int, exclude map[int]string) (dataServers []string) {
	candidates := make([]string, 0)
	reverseExcludeMap := make(map[string]int)
	for id, address := range exclude {
		reverseExcludeMap[address] = id
	}
	servers := GetDataServers()
	for i := range servers {
		server := servers[i]
		_, exclude := reverseExcludeMap[server]
		if !exclude {
			candidates = append(candidates, server)
		}
	}
	length := len(candidates)
	if length < n {
		return
	}
	part := rand.Perm(length)
	for i := 0; i < n; i++ {
		dataServers = append(dataServers, candidates[part[i]])
	}
	return
}
