# go-design-patterns

This is an exercise in the exploration of different design patterns using the Golang language.


Different formats of telematics data have been imagined, and we're supposed to extract the **odometer** field out of them.
The formats are classified on the basis of the different models of vehicles, which are as follows:


## Model A : 
```
{
    model: A,
    data: {
        odometer:
    }
}
```

## Model B : 
```
{
    model: B,
    data: {
        speed: {
            odometer:
        }
    }
}
```

## Model C: 
```
{
    model: C,
    data: [
        {
            spnId: 
            pgnId:
            spnName: odometer
            value:
        }
    ]
}
```