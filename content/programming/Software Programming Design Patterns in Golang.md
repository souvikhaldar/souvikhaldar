---
title: Software Programming Design Patterns in Golang
date: 2022-11-18
draft: false
---

# Introduction
Code can be written in any way till the time it works, but if you write them following certain time-tested design patterns, it can lessen a lot of future headaches that might otherwise happen while scaling it further in due course of time. A well-designed software can be 
1. Extensible, so extending functionalities will come at least cost of modifications.  
2. Simple, so reading the code can give idea of what it is trying to do easily.  
3. Abstracted, so not everyone needs to dig deep into the code in order to be able to use it.  

# Design patterns
## Builder Pattern
In builder pattern, a complex object is created using Setters of every field to set the value of the fields, then calling a `Build` method to return an object is all fields meet required condition or are ready to be used. 
```
type Person struct{
	firstName string
	lastName string
}
func (p *Person)SetFirstName(name string){
	// Check for required conditions for using first name here
	p.firstName = name
}

func (p *Person)SetLastName(name string){
	p.lastName = name
}
func (p *Person)Build()(*Person,error){
	// Check for required conditions for building person object here
	return &Person{
		fistName: p.firstName,
		lastName: p.lastName,
	}
}
```
It allows you to each field anytime wanted, and its own set of conditions to check.
Also, the build process and check for errors while building the entire object, which might require calling APIs, calling databases, etc in the process. 

## Singleton Pattern
Sometimes it is required that we instantiate an object once in the lifetime of running of the program, like object that stores connection to database and many more such cases. Hence, this pattern strictly makes sure that we do not instantiate it more than once time even if called to do so.

```
var db *DB.Cursor
func NewDB()*DB.Cursor{
	sync.Once(func(){
		db = createConnectionToDatabase()
	})
	return db
}
```

In the above case, we could have checked if `db == nil` and then only create the object, but in doing it would not have been thread safe. Use `sync.Once` makes sure that the function passed to it runs only once.

## Observer Pattern
Sometimes, we need to react to change of state of objects in the code. We want to observe the change in the state of the object and whenever it changes, we want a certain reaction to it. 
Let's try to see how this pattern can be achieved.

```
-----Observer.go----------
type Observer interface{
	onChange() // upon observing change this should be triggered
}

type Observe struct{}
func (o *Observe)onChange(){
	// do something
}

----- data.go ---------
type Observed interface{
	registerObserver(o Observer)
	change()
	notifyAll()
}
type Data struct{
	observers []Observer
}
func (d *Data)registerObserver(o Observer){
	d.observers = append(d.observers,o)
}
func (d *Data)Change(){
	// change happened to the observed data
	d.notifyAll()
}

func(d *Data)notifyAll(){
	for _, o := range d.observers{
		o.onChange()
	}
}
```

## Repository Pattern
Respository Pattern is a design in which we abstract implementation details by using one interface per package called `Repository` and then implementing it by different use cases. This makes changing the underlaying technology very easy since we are not tightly coupled with the implementation details.
```
---------db.go-------------
type User interface{
	addUser(userName string)
}

---------postgres.go--------
type Db struct{
	db *db.Conn
}
func(d *Db)addUser(username string){
	// add user to db
}

--------- main.go ----------
type server struct{
	userRepo db.User
}
func main(){
	s := server{
		userRepo: postgres.DB, // here we are using postgres but can easily swap with any other database
	}
	
	s.addUser("souvik")
}

```
## Conclusion
Hope you've got an idea of some of the popular design patterns that you can use to improve the quality of overall codebase. This is definitely not the exhaustive list. Please feel free to provide feedbacks to me on any medium of choice, links to which are provided in the homepage of my website. Happy coding! 

## References
1. https://www.linkedin.com/learning-login/share?account=92962490&forceAccount=false&redirect=https%3A%2F%2Fwww.linkedin.com%2Flearning%2Fgo-design-patterns%3Ftrk%3Dshare_ent_url%26shareId%3DshGy9CWZT8SPXxYjGD1AEw%253D%253D