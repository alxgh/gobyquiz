/*
Good! now that you ran your first program you get to know pacakges!

In short a package is some piece of related code together somewhere which are exposed to you for usage.

We have the Golang's standard packages such as math and net (which probably by the name you can guess what they do) and also packages that others have published on the internet usually through a version manager like git.

If you have used some other languages like Javascript or Python it is good to know that there is no single source for pacakges for Golang, You can create a Github repo and that will be a package that others can use in their programs!

Back to the question: Here we have used the "fmt" packages which is used for formatting and priting to console.
To import a package simply use import keyword followed byu pacakge name surrouned in double quotes.

Also you can group multiple packages in different lines between paranthesis after the import keyword!
*/
package main

func main() {
	fmt.Println("Hey there!")
}
