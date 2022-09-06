# REPO ANTIGUO
Ya no hace falta usar esta base para armar un scraper. en la nueva versión se puede modificar el dominio del sitio directamente con un archivo yaml. ver
[https://github.com/anongolico/zcrapper]


# Base

Módulo para modelar la información que devuelve la página.\
**Ojo:** este repo no sirve para descargar nada de internet, solamente modela la info. Para ver una implementación práctica, pasate por este otro lado:\
- [zcrapper](https://github.com/anongolico/zcrapper)



### Para usarlo en tu proyecto:
```
go get -u -v github.com/anongolico/base
```

***

### ¿Qué información puedo obtener?

Actualmente soporta:

#### Rouz
1) Id (e.g. *KKOIO8Q8UMFSUOQDS6DI*)
2) Título
3) Archivo de portada

#### Comentario
1) Id (e.g. *JMYC1LKZ*)
2) Contenido (e.g. *'y nene down??'*)
3) Fecha de creación
4) Archivo adjunto

### Eres un inútil, OP, te falta agregar X campo
Clona el repo
```
git clone https://github.com/anongolico/base.git
```
Y agrega los campos necesarios en cada struct. En este mismo repo hay un archivo `ejemplo.json` para ver la estructura de la info que quieras parsear.


## Clones
Los clones del [repo](https://github.com/LaDevWendy/rouzer3.0) original de rouzer siguen la misma estructura. Por ejemplo, [Ufftopia](https://ufftopia.net) es perfectamente compatible con este módulo. No más hay que cambiar las variables globales con las URL y listo, haz creado una base para parsear posts de ufftopia.
