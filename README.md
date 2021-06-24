# Apache Benchmark

Autor: Sergi Gomà Cruells



## Anàlisi dels paràmetres de Apache Benchmark

*Totes les proves s'han realitzat amb un processador i7-7700k de quatre cores a 4,82 GHz.*



### Paràmetre -c

El paràmetre -c es fa servir per indicar quantes transaccions es fan a l'hora, és a dir, és un paràmetre per configurar la concurrència. Fent servir la comanda `docker run --rm jordi/ab -c 10 -n 1000 http://www.docker.com/`, faré 1.000 transaccions a la web de docker, amb diversos valors de concurrència.

Aquests són els resultats:

| -c   | TPS (#/s) | Latency (ms) | CPU (%) |
| ---- | --------- | :----------- | ------- |
| 1    | 32.94     | 30           | 1.3     |
| 2    | 65.77     | 30           | 2.6     |
| 3    | 99.92     | 30           | 3.7     |
| 4    | 132.44    | 30           | 5       |
| 5    | 163.3     | 30           | 6       |
| 6    | 196.33    | 30           | 7.5     |
| 10   | 290.3     | 34           | 9.6     |
| 20   | 337.36    | 58           | 11.3    |
| 50   | 342.54    | 140          | 11      |
| 80   | 340.29    | 225          | 10.8    |
| 100  | 346.83    | 278          | 11.3    |
| 150  | 351.72    | 404          | 11.1    |
| 200  | 343.41    | 543          | 10.7    |
| 250  | 333.38    | 662          | 10.9    |
| 400  | 338.61    | 731          | 11.3    |
| 500  | 347.11    | 910          | 11      |

![Texto alternativo](tpslatencychart.png)

Els resultats mostren que a mesura que s'incrementa la concurrència, les transaccions per segon augmenten fins que arriben a unes 340 aproximadament i nivell de concurrència 20, on s'estabilitzen. A partir d'aquí, un valor més alt de concurrència només introduirà més latència, sense millorar el temps d'execució. Respecte al percentatge d'ús de la CPU, s'incrementa fins a una utilització al voltant del 10%, que coincideix amb un nivell de concurrència 20.



### Paràmetre -k 

Aquest paràmetre serveix per habilitar la funció KeepAlive d'HTTP, que permet mantindre la connexió mentre encara estem fent transaccions al servidor, enlloc de reiniciar-la cada cop. Això comporta una millora de la eficiència, és a dir, menor latència.
