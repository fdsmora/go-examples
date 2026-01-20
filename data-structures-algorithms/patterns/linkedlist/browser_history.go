package linkedlist

// LeetCode problem 1472. Design Browser History
// https://leetcode.com/problems/design-browser-history/
// Medium
type (
	Page struct {
		url                string
		nextPage, prevPage *Page
	}

	BrowserHistory struct {
		currentPage *Page
	}
)

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		currentPage: &Page{url: homepage},
	}
}

func (this *BrowserHistory) Visit(url string) {
	//fmt.Printf("visiting. Curr: %s, next: %s\n", this.currentPage.url, url)
	nextPage := &Page{url: url}
	nextPage.url = url
	nextPage.prevPage = this.currentPage
	this.currentPage.nextPage = nextPage
	this.currentPage = this.currentPage.nextPage
	// fmt.Printf("visiting. Now Curr: %s \n", this.currentPage.url)

}

/*
lc<->g<->f-x-y

	<->l

c
*/
func (this *BrowserHistory) Back(steps int) string {
	i := 0
	for i < steps && this.currentPage.prevPage != nil {
		//fmt.Printf("Back. it: %d, curr: %s, prev: %s\n", i, this.currentPage.url, this.currentPage.prevPage.url)
		this.currentPage = this.currentPage.prevPage
		i++
	}
	return this.currentPage.url
}

func (this *BrowserHistory) Forward(steps int) string {
	i := 0
	for i < steps && this.currentPage.nextPage != nil {
		// fmt.Printf("it: %d, curr: %s, prev: %s\n", i, this.currentPage.url, this.currentPage.nextPage.url)
		this.currentPage = this.currentPage.nextPage
		i++
	}
	return this.currentPage.url
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
