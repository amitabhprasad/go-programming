### Pointers and methodset
- If interface is implemented by receiver that is on the pointer
    then it can be called only by the pointer value
- If interface is impleted by type 
    then it can be called either by the type or the pointer
Receiver        Values
*type   -->     *type
type    -->     Type, *type