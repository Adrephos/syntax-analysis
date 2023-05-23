import string
import random as ran
from ordered_set import OrderedSet


def eliminar_recursividad_izquierda(gramatica_ini):
    # imprimir G inicial
    print("Gramatica inicial: ",gramatica_ini)
    gramatica = gramatica_ini
    nuevas_producciones = []
    no_terminales = OrderedSet()
    cont_i=1
    no_terminales_ini = []
    no_terminales_ini = list(NoTerminalesIniciales(gramatica_ini))
    print("no terminales iniciales", no_terminales_ini)
    print("-------- INICIO------------")
    # lista para verificar j
    no_term_auxiliares = OrderedSet()
    g_auxiliar=[]
    
    
    #paso por producciones
    for produccion in gramatica:
        # separo porducciones de no terminal
        no_terminal, opciones = produccion.split(" -> ")
        # separo cada produccion
        opciones = opciones.split(" | ")
        print("produ inicial:", produccion)
        # Recursivas y no recursivas por izquierda (γ β)
        recursivas = []
        no_recursivas = []

        # ciclo de j que verifica Ai -> Aj prods
        if cont_i >1:
            cont_j=0
            nuevas_producciones.append(produccion)

            for j in  range(0, cont_i-1):

                producto = nuevas_producciones[-1]
                no_terminal, opciones = producto.split(" -> ")
                opciones = opciones.split(" | ")

                no_term_auxiliares = no_terminales
                lista_no_term_aux = list(no_term_auxiliares)

                referencia = nuevas_producciones[j]
                no_terminal_aux, opciones_aux = referencia.split(" -> ")
                opciones_aux = opciones_aux.split(" | ")
                
                producciones_Ai=[] # gamma o parte de las prod sin recursividad por izquierda
                deltas=[] # delta o parte de las prod con recursividad por izquierda
                ref = None

                cont = 0
                for opcion in opciones:
                    if any(opcion.startswith(elem) for elem in no_terminales_ini[j]): # agrego delta
                        deltas.append(opcion[1:])

                        ref = True
                        cont +=1
                    else:
                        producciones_Ai.append(opcion)
                

                if ref is True:
                    cont_j+=1
                    producciones_Aj= conseguirProducciones(g_auxiliar, no_terminal_aux)
                    print("hay rec hacia atras en", cont_i, cont_j)
                    nuevas_producciones.pop()
                    produccion_reemplazo = crearProduccionReemplazo(producciones_Ai, deltas, producciones_Aj, no_terminal)
                    nuevas_producciones.append(produccion_reemplazo)
                    print("g parcial ", nuevas_producciones, " en i,j: ", cont_i,",", j+1 )
                else:
                    cont_j +=1
                    print("No hay rec hacia atras en", cont_i, cont_j)

            if cont == 0:
                print("No hay rec hacia atras en", cont_i)
            

        else:
            nuevas_producciones.append(produccion)
        print("<<<<<<<<<<<<<<<<<")
        # fin del ciclo Ai -> Aj
        production = nuevas_producciones[-1]
        print("produccion",production)

        no_terminal, opciones = production.split(" -> ")
        opciones = opciones.split(" | ")
        no_terminales.add(no_terminal)
        for opcion in opciones:
            if opcion.startswith(no_terminal):
                recursivas.append(opcion[1:])
            else:
                no_recursivas.append(opcion)
        
        # ---------------- Eliminating Direct Left Recursion
        if len(recursivas) > 0: # en caso de Existir β
            print("hay recursividad directa")
            nuevas_producciones.pop()
            nuevo_no_terminal = randomChar(no_terminales, gramatica_ini)
            no_terminales.add(nuevo_no_terminal)

            
            if('ε') in no_recursivas:
                no_recursivas.remove('ε')
                no_recursivas.append('ε')
            
            # Para todas las opciones γ les agrego el no terminal Auxiliar
            var_reemplazo = no_terminal + " -> " + " | ".join([elem + nuevo_no_terminal if elem != "ε" else nuevo_no_terminal for elem in no_recursivas])
            nuevas_producciones.append(var_reemplazo)
            
            # opciones recursivas van a produccion de no terminal auxiliar
            var_produccion = nuevo_no_terminal + " -> " + " | ".join([elem + nuevo_no_terminal for elem in recursivas]) + " | ε"
            nuevas_producciones.append(var_produccion)
        else:
            print("no hay rec directa")
        
        
        
        g_auxiliar = nuevas_producciones
        print("g parcial: ",nuevas_producciones, "en ", cont_i)
        print("---------------------------")
        cont_i +=1
    
    return nuevas_producciones, no_terminales

# funciones para generar nuevos no terminales distintos a los ecxistentes
def randomChar(no_terminales, g):
    random_posible = [char for char in string.ascii_letters.upper() if char not in (no_terminales) and char not in (NoTerminalesIniciales(g)) ]
    if random_posible:
        random = ran.choice(random_posible)
        return random
    else:
        return ("?")
    
def conseguirProducciones(gramatica, no_terminal):
    resultados = []
    for produccion in gramatica:
        if produccion.startswith(no_terminal + ' ->'):
            no_term, produccion_separada = produccion.split(' -> ')
            producciones = produccion_separada.split(' | ')
            resultados.extend(producciones)
    return resultados

def crearProduccionReemplazo(Ai, deltas, Aj, no_terminal_aux):
    cambio_prod_aux=[]
    auxiliar = []
    reemplazoProducciones = [x+y for x in Aj for y in deltas]
    producciones = Ai + reemplazoProducciones
    pruducto_Ai= no_terminal_aux + " -> " + " | ".join(producciones)
    
    no_terminal, opciones = pruducto_Ai.split(" -> ")
    # separo cada produccion
    opciones = opciones.split(" | ")
    for opcion in opciones:
        opcion = eliminarEpsilon(opcion)
        auxiliar.append(opcion)
    pruducto_Ai= no_terminal_aux + " -> " + " | ".join(auxiliar)

    cambio_prod_aux.append(pruducto_Ai)
    print("producto creado final: ", pruducto_Ai)
    
    produccion_reemplazo = " ".join(cambio_prod_aux)
    return (produccion_reemplazo)

def pedirGramatica():
    print("A continuación ingrese la Gramática que desee procesar, recuerde que la gramática debe ser de la forma A ->  α |  ßB <enter> B-> ε <enter> donde:")
    print("A es no terminal y todas sus producciones van después de -> separadas por |, además α y ß se componen por símbolos terminales o No terminales")
    print("Ingrese la Gramatica que desea procesar: ")
    gramatica = []
    while True:
        rule = input()
        if rule == "$": 
            break
        gramatica.append(rule)

    for index, produccion in enumerate(gramatica):
        gramatica[index] = produccion.strip()
        print(gramatica[index])
    return (gramatica)

def NoTerminalesIniciales (g):
    resultado= OrderedSet()
    for produccion in g:
        no_term, residuo= produccion.split(" -> ")
        resultado.add(no_term)

    return resultado

def eliminarEpsilon (opcion):
    if opcion == "ε":
        aux = opcion        
    elif "ε" in opcion:
        aux =  opcion.replace("ε", "")
    else:
        aux = opcion
    return aux
# Ejemplo de gramática
# donde A es no terminal y todas sus producciones van despues de -> separadas por |
""" 
# ejemplos de g

gramatica_ini =[
    "S -> Aa | b",
    "A -> Ac | Sd | ε"
]
# S -> Aa | b,A -> Ac | Sd | ε

gramatica_ini = [
    "T -> TbT | V",
    "V -> cV | c"
]
# T -> TbT | V,V -> cV | c

gramatica_ini = [ # no tiene lr
    "S -> AS | A",
    "A -> 0A | 1B | 1",
    "B -> 0B | 0"
]
# S -> AS | A,A -> 0A | 1B | 1,B -> 0B | 0

gramatica_ini = [
    "S -> S0 | T",
    "T -> S1 | T1 | 1"
]
# S -> S0 | T,T -> S1 | T1 | 1

gramatica_ini = [
    "A -> BC | a",
    "B -> CA | Ab",
    "C -> AB | CC | a"
]
# A -> BC | a,B -> CA | Ab,C -> AB | CC | a

gramatica_ini = [
    "S -> Aa | Bc",
    "A -> Dd | a",
    "B -> d | ε",
    "D -> Sa"
]
# S -> Aa | Bc,A -> Dd | a,B -> d | ε,D -> Sa

"""

# pedir gramatica
gramatica_ini = pedirGramatica()

# Eliminar recursividad por la izquierda
resultado, no_terminales = eliminar_recursividad_izquierda(gramatica_ini)

# Imprimir las nuevas producciones
print("Gramatica sin Recursividad por Izquierda: -----------")
for produccion in resultado:
    print(produccion)

# Imprimir los no terminales
print("Conjunto de No terminales: ", no_terminales)
