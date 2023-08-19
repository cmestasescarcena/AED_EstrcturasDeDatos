#include <iostream>
#include <vector>
#include <algorithm>
#include <stdlib.h>
#include <time.h>
#include <chrono>
#include <random> 

using namespace std;
 
// Estructura de datos para almacenar un borde de graph
struct Edge {
    int src, dest, weight;
};
 
typedef pair<int, int> Pair;
 
// Una clase para representar un objeto graph
class Graph
{
public:
    // un vector de vectores de Pares para representar una lista de adyacencia
    vector<vector<Pair>> adjList;
 
    // Constructor de graph
    Graph(vector<Edge> const &edges, int n)
    {
        // cambiar el tamaño del vector para contener `n` elementos de tipo `vector<int>`
        adjList.resize(n);
 
        // agrega bordes al grafo dirigido
        for (auto &edge: edges)
        {
            int src = edge.src;
            int dest = edge.dest;
            int weight = edge.weight;
 
            // insertar al final
            adjList[src].push_back(make_pair(dest, weight));
 
            // elimine el comentario del siguiente código para el graph no dirigido
            adjList[dest].push_back(make_pair(src, weight));
        }
    }
};
 
// Función para imprimir la representación de la lista de adyacencia de un graph
void printGraph(Graph const &graph, int n)
{
    for (int i = 0; i < n; i++)
    {
        // Función para imprimir todos los vértices vecinos de un vértice dado
        for (Pair v: graph.adjList[i]) {
            cout << "(" << i << ", " << v.first << ", " << v.second << ") ";
        }
        cout << endl;
    }
}


//*******************************************************

class Route
{
  public:
    vector<int> Nodes;
    int Totaldistance;
  Route()
  {
    vector<int> Nodes;
    Totaldistance = 0;
  };
};


Route Algorithm(int origin, Graph graph, int seed)
{
  Route solution;

  for (int i=0; i<graph.adjList.size(); i++)
  {
    solution.Nodes.push_back(i);
  }

  shuffle(solution.Nodes.begin(), solution.Nodes.end(),default_random_engine(seed));


  for (int i=0; i<solution.Nodes.size(); i++)
  {
    if(solution.Nodes.at(i) == origin)
    {
      solution.Nodes.erase(solution.Nodes.begin()+i);
    }
  }

  solution.Nodes.insert(solution.Nodes.begin(), origin);
  solution.Nodes.push_back(origin);

  for(int i=0; i < solution.Nodes.size()-1; i++)
  {
    int Fnode = solution.Nodes.at(i);

    for(int j=0; j < graph.adjList.size()-1; j++)
    {
      if (graph.adjList.at(Fnode).at(j).first == solution.Nodes.at(i+1))
      {
        solution.Totaldistance=solution.Totaldistance + graph.adjList.at(Fnode).at(j).second;
      }
    }
  }

  return solution;
}
