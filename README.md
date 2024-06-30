# Program Checker

Build exec : 

```sh
go build -o checker main.go
```

Différents test : 

```sh
./checker "3 2 1 0"
```

The program gonna ask you to enter instructions :

```sh
pierrecaboor@macbook-pro checker % ./checker "3 2 1 0"        
Enter instruction (or press [Enter] again to finish): sa
Executing instruction: sa
Before - A: [3 2 1 0], B: []
After - A: [3 2 0 1], B: []
Enter instruction (or press [Enter] again to finish): rrr
Executing instruction: rrr
Before - A: [3 2 0 1], B: []
After - A: [1 3 2 0], B: []
Enter instruction (or press [Enter] again to finish): 
KO
```

You can follow différents steps and see the result : `KO` or `OK`
