## The Adapter Pattern -- Part I: a basic understanding

Co-authored-by: Graham Holtslander <gholtslander@vendasta.com>


### Preface

Back in the late 2000s when Android became a popular smartphone OS, I was among those tech geeks
who wanted to use an Android phone badly. However, back then the Android phones were all
imported to mainland China and they were expensive. Although I finally convinced my parents to buy me a Samsung
Galaxy III, I found a problem charging the phone; its charger wasn't compatible with a standard Chinese
power outlet. OMG! I panicked when the phone battery was dying and right at the moment my roommate told me that I might
need to buy an adapter. Okay, so after withdrawing a couple grand RMB
(which was amount to my 4-month living expense back then) from my bank account, I spent another
20 to get an adapter. I know what you are about to say about me, I'm cheap :facepalm:
Well, the good thing about that experience is I get to know what an adapter is and that's exactly what
we are talking about today in this post.


### Definition of Adapter

So from the above story you may already have an idea of what an adapter is in terms of object-oriented programming.
According to Gamma et al., an adapter converts the interface of a class (the *adaptee*) into another interface clients
expect. Adapter is also known as **wrapper** and can be mixed with other dessign patterns;
**decorator**, **facade**, **bridge** and **proxy** (hopefully chapters about these ones will come soon).
In the previous example, the Galaxy III charger from its original box is the client who wants to
use a Chinese power outlet and the Chinese power outlet becomes the adaptee.
Of course the AC adapter is the adapter.


### Motivation

Now let's get a little deeper, why do we need an adapter in OOP?
The idea behind this design pattern is that we don't want to change any code of the existing client or adaptee,
especially when they've been well tested and maintained. In the previous example, we may not want to use
a different charger because it may not provide the same current to the battery
(or may even blow the phone entirely because we know Samsung battery can be a public safety concern);
neither do we want to create
a new power outlet just to charge an exotic phone because it is expensive and barely reusable to other device chargers.
In this regard, the most cost-efficient solution is write an adapter.


### Code Example

There are numerous uses cases and examples all over the places but I find
[this one](https://www.youtube.com/watch?v=qG286LQM6BU)
particularly comprehensible. However, I'd like to use `go` code to elaborate this example.
Assume in a game project we've already had an `enemyAttacker` ready and an `enemyTank` implements
this interface. Also, in addition to tanks, we have another `enemyRobot` class ready.

```go
// enemyAttacker interface defines what an attacker does in general
type enemyAttacker interface {
	fireWeapons()
	move()
	assignDriver(driverName string)
}

func run(a enemyAttacker, driverName string) {
	a.assignDriver(driverName)
	a.fireWeapons()
	a.move()
}
```

```go
// enemyTank implements enemyAttacker's methods
type enemyTank struct {
	damage int
	moveSpace int
}

func NewEnemyTank() *enemyTank {
	fmt.Println("A new enemy tank has been created")
	return &enemyTank{
		damage:    rand.Intn(20) + 1,
		moveSpace: rand.Intn(10) + 1,
	}
}

func (t enemyTank) fireWeapons() {
	fmt.Println("Cause", t.damage, "damages")
}

func (t enemyTank) move() {
	fmt.Println("Move", t.moveSpace, "spaces")
}

func (t enemyTank) assignDriver(driverName string) {
	fmt.Println("Assign to driver", driverName)
}
```

```go
// enemyRobot has a set of different methods than enemyAttacker
type enemyRobot struct {
	smashDamage int
	walkDistance int
}

// A robot causes less damage and moves slower than a tank
func NewEnemyRobot() *enemyRobot {
	fmt.Println("A new enemy robot has been created")
	return &enemyRobot{
		smashDamage:  rand.Intn(10) + 1,
		walkDistance: rand.Intn(5) + 1,
	}
}

func (r enemyRobot) smashWithHands() {
	fmt.Println("Smash cause", r.smashDamage, "damages")
}

func (r enemyRobot) walk() {
	fmt.Println("Walk in", r.walkDistance, "spaces")
}

func (r enemyRobot) marchWithDriver(driverName string) {
	fmt.Println("Stop tramping, marching with", driverName)
}
```

As we can see, `enemyRobot` has a set of different methods that makes it incompatible with `enemyAttacker`.
Also, in `go`, it throws an error upon calling
`run(robot, "A poor dirver")`, where `robot` is an instance of `enemyRobot`. In this case,
we want to write a class (in `go` it's rather a struct) to adapt the `enemyAttacker` to use `enemyRobot`.
We then write a `robotAdapter`:

```go
// robotAdapter defines an adapter for enemyAttacker to use enemyRobot
type robotAdapter struct {
	robot *enemyRobot
}

func NewRobotAdapter(r *enemyRobot) *robotAdapter {
	fmt.Println("A new enemy robot adapter has been created")
	return &robotAdapter{
		robot: r,
	}
}

func (a robotAdapter) fireWeapons() {
	a.robot.smashWithHands()
}

func (a robotAdapter) move() {
	a.robot.walk()
}

func (a robotAdapter) assignDriver(driverName string) {
	a.robot.marchWithDriver(driverName)
}
```

Not quite complicated eh?
Finally in our `main` function we can test this example.

```go
func main() {
	// enemy tank
	tank := NewEnemyTank()
	run(tank, "Nathan")

	// enemy robot
	robot := NewEnemyRobot()
	robot.marchWithDriver("Graham")
	robot.smashWithHands()
	robot.walk()
	// the following commented out lines will have compile errors because of incompatibility
	// run(robot, "Graham") // DO NOT do this
	// However, the following lines will pass and work just fine
	// enemy adapter
	adapter := NewRobotAdapter(robot)
	run(adapter, "Donald")
}
```

The following shows the result of running `main`
```shell script
rocky:adapter_pattern nyang$ go run adapter_example.go
A new enemy tank has been created
Assign to driver Nathan
Cause 2 damages
Move 8 spaces
A new enemy robot has been created
Stop tramping, marching with Graham
Smash cause 8 damages
Walk in 5 spaces
A new enemy robot adapter has been created
Stop tramping, marching with Donald
Smash cause 8 damages
Walk in 5 spaces
```

### Summary

There are 4 essential parts in the adapter design patter:

1. **Target**
    - defines the domain-specific interface that **Client** uses
1. **Client**
    - collaborates with objects conforming to the **Target** interface
1. **Adaptee**
    - defines an existing interface that needs adapting
1. **Adapter**
    - adapts the interface of **Adaptee** to the **Target** interface

To follow Graham's writing style, I'd like to keep this post short but readable.
Hopefully this brief post helps you to capture the basics of the adapter design pattern.
There are, for sure, more in-depth illustration of the adapter pattern and I (or someone else who would like to)
will wrap up the content in Part II to discuss more about adapters.