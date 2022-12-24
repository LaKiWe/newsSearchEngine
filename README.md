
Default Addr:localhost:9134 

## Storage design


## Operate design
### Backend

Engine{
    index\doc DB ptrs,
    operation(resquest <-> db) implement),
}

DB{
    disk stroage{level db}:actually storage doc,words and index,
    mem storage{mainly id-score tree,runtime cache}:io cache, tmp scores caculate, tree stroage),
}

### Frontend

