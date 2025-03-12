package main

import "fmt"

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	friends_set := make(map[string]bool)

	user_friends := make(map[string]bool)
	for _, friend := range friendships[username] {
		user_friends[friend] = true
	}
	for _, friend := range friendships[username] {
		for _, f := range friendships[friend] {
			if f != username && !user_friends[f] {
				friends_set[f] = true
			}
		}
	}

	var suggested_friends []string
	for friend := range friends_set {
		suggested_friends = append(suggested_friends, friend)
	}

	if len(suggested_friends) == 0 {
		return nil
	}

	return suggested_friends
}

func main() {
	friendships := map[string][]string{
		"Alice":   {"Bob", "Charlie"},
		"Bob":     {"Alice", "Charlie", "David"},
		"Charlie": {"Alice", "Bob", "David", "Eve"},
		"David":   {"Bob", "Charlie"},
		"Eve":     {"Charlie"},
	}

	suggestedFriends := findSuggestedFriends("Alice", friendships)
	fmt.Println(suggestedFriends) // i'm expecting -> [David Eve]
}
