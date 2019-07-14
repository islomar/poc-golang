# Ejercicio para practicar HTTP requests

Enunciado: https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/59283627/


Escoge una API JSON que sea de tu agrado (ahí van algunas propuestas):
* https://pokeapi.co 
* https://swapi.co 
* https://api.github.com 
* https://punkapi.com

Escribe un pequeño programa que:

 - Haga una petición GET a alguno de los endpoints de la API.
 - Recoja los datos (JSON) y los deserialice en memoria.
 - Serialice los datos anteriores en formato CSV
 - Guarde el fichero CSV.
 
 
Algunos de los paquetes que vas a necesitar:

 - net/http
 - encoding/json
 - encoding/csv
 - bufio
 - io/ioutil
 
 
Algunas funcionalidades extra que podrías incorporar:

 - Que sea un CLI con flags.
 - Que le puedas especificar el nombre del fichero de salida.
 - Que le puedas dar un endpoint de entrada (high-level!).