# go-examples
Mis propios ejemplos para aprender go

## Notas

- Tipos que se pasan por referencia por default:
  - map
  - slices
  - pointers

- Los tipos de datos `char` se conocen como `runes` y es un alias para `int32`.

- `make`: The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it. 

- Diferencia entre `make` y `new`: `new` aloja memoria para una variable del tipo especificado y retorna un apuntador a la memoria alojada. 
