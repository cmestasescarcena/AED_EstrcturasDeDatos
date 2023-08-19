#include <iostream>
#include <math.h>
#include <stdlib.h>
#include <time.h>
#include <sys/time.h>
#include "graph.hpp"

using namespace std;

long obtenerTiempo();

int main()
{
  
  srand(time(NULL));

    // vector de los bordes del graph según el diagrama anterior.
    // Tenga en cuenta que el vector de inicialización en el siguiente formato
    vector<Edge> edges =
    {
        // (x, y, w) —> arista de `x` a `y` con peso `w`
        {0, 1, 234}, {0, 2, 4234}, {0, 3, 234}, {0, 4, 423}, {0, 5, 423}, {0, 6, 234}, {0, 7, 564}, {0, 8, 465}, {0, 9, 546}, {0, 10, 12}, {0, 11, 1233}, {0, 12, 1234},
        {1, 2, 345}, {1, 3, 567}, {1, 4, 867}, {1, 5, 123}, {1, 6, 234}, {1, 7, 7456}, {1, 8, 234}, {1, 9, 235}, {1, 10, 456}, {1, 11, 453}, {1, 12, 345},
        {2, 3, 645}, {2, 4, 456}, {2, 5, 769}, {2, 6, 124}, {2, 7, 513}, {2, 8, 645}, {2, 9, 98}, {2, 10, 456}, {2, 11, 456}, {2, 12, 834},
        {3, 4, 678}, {3, 5, 567}, {3, 6, 456}, {3, 7, 34}, {3, 8, 345}, {3, 9, 564}, {3, 10, 234}, {3, 11, 234}, {3, 12, 645},
        {4, 5, 234}, {4, 6, 24}, {4, 7, 345}, {4, 8, 678}, {4, 9, 534}, {4, 10, 132}, {4, 11, 879}, {4, 12, 6542},
        {5, 6, 345}, {5, 7, 345}, {5, 8, 345}, {5, 9, 543}, {5, 10, 564}, {5, 11, 345}, {5, 12, 2346},
        {6, 7, 354}, {6, 8, 675}, {6, 9, 453}, {6, 10, 123}, {6, 11, 678}, {6, 12, 834},
        {7, 8, 234}, {7, 9, 345}, {7, 10, 453}, {7, 11, 456}, {7, 12, 345},
        {8, 9, 342}, {8, 10, 345}, {8, 11, 345}, {8, 12, 456},
        {9, 10, 453}, {9, 11, 645}, {9, 12, 756},
        {10, 11, 234}, {10, 12, 345},
        {11, 12, 456}
    };
 
    // número total de nodos en el graph (etiquetados de 0 a 5)
    int n = 13;
 
    // construir grafo
    Graph graph(edges, n);

    // imprime la representación de la lista de adyacencia de un graph
    // printGraph(graph, n);
    
    cout<<endl;
    cout<<endl;

    for(int i=0; i<50; i++)
    {
      int _a = rand();

      long inicio = obtenerTiempo();
      double tiempoEnSegundos;
      long tiempoEnMicrosegundos;

      for(int j=0; j<Algorithm(0, graph, _a).Nodes.size(); j++)
      {
        cout<<Algorithm(0, graph, _a).Nodes.at(j)<<" ";
      }

      long final = obtenerTiempo();
      tiempoEnMicrosegundos = final - inicio;
      tiempoEnSegundos = tiempoEnMicrosegundos * pow(10, -6);
      cout<<" -> Cost: "<<Algorithm(0, graph, _a).Totaldistance<<" -> Time: "<<tiempoEnSegundos;
     
      cout<<endl;
    }

    return 0;

}

long obtenerTiempo(){
  struct timeval inicio;
  gettimeofday(&inicio, NULL);
  return inicio.tv_sec*1000000+inicio.tv_usec;
}
