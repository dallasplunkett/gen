# gen
CLI tool for generating CSV files with synthetic data.

## Usage
gen "filename.csv" c1:norm(mu,std, s, e) c2:lin(m, b, s, e) -r 200

## todo / learn
* parse parameters
* parse


## ideas
* instead of operating on both integers and floats, just start with floats and maybe add integer support later

```
map[
    column: function{
        name: "normal",
        params: [1.0, 23.4, 9.5],
    }
]
```