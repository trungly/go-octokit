package octokit

var (
	CurrentFollowerUrl  = Hyperlink("user/followers")
	FollowerUrl         = Hyperlink("users/{user}/followers")
	CurrentFollowingUrl = Hyperlink("user/following{/target}")
	FollowingUrl        = Hyperlink("users/{user}/following{/target}")
)

// Create a FollowersService
func (c *Client) Followers() (followers *FollowersService) {
	followers = &FollowersService{client: c}
	return
}

// A service to return user followers
type FollowersService struct {
	client *Client
}

// Get a list of followers for the user
func (f *FollowersService) All(uri *Hyperlink, uriParams M) (followers []User, result *Result) {
	url, err := ExpandWithDefault(uri, &CurrentFollowerUrl, uriParams)
	if err != nil {
		return nil, &Result{Err: err}
	}

	result = f.client.get(url, &followers)
	return
}

// Checks if a user is following a target user
func (f *FollowersService) Check(uri *Hyperlink, uriParams M) (success bool, result *Result) {
	url, err := ExpandWithDefault(uri, &CurrentFollowingUrl, uriParams)
	if err != nil {
		return false, &Result{Err: err}
	}

	result = f.client.get(url, nil)
	success = (result.Response.StatusCode == 204)
	return
}

// Follows a target user
func (f *FollowersService) Follow(uri *Hyperlink, uriParams M) (success bool, result *Result) {
	url, err := ExpandWithDefault(uri, &CurrentFollowingUrl, uriParams)
	if err != nil {
		return false, &Result{Err: err}
	}

	result = f.client.put(url, nil, nil)
	success = (result.Response.StatusCode == 204)
	return
}

// Unfollows a target user
func (f *FollowersService) Unfollow(uri *Hyperlink, uriParams M) (success bool, result *Result) {
	url, err := ExpandWithDefault(uri, &CurrentFollowingUrl, uriParams)
	if err != nil {
		return false, &Result{Err: err}
	}

	result = f.client.delete(url, nil, nil)
	success = (result.Response.StatusCode == 204)
	return
}
