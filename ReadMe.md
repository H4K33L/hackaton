# Whelcome to Generator 5ém monster project !

The purpose of the project is to generate information for monster of the 5em Edition.
(For more information about 5em Edition see https://www.aidedd.org/dnd-5/)
The monster information generated are : Name, caracteristics, resistance, vulnerability, imunity ...
All information exept attack and sepcials deals.
The users use a localhost web site to generate monster, the users can chose numbers of monster, the 
type of monster and the ID (Indice of dificulties). The web site also print all the monster generated,
when user click on the name of a monster it print all information about the monster bellow the monsters 
list.

# ------------------------------------------------------------------------------------------------------- #

# Project status :

The project still in Beta, it whorks in is main purpose, generate monster profiles. But somme features can
and need to be added to the project :

 - Monster Group generation (only on the web site there is a monster group generation)
 - Encounter table generation (only on the web site there is a Enconter table generation)
 - improve user experience on the web site

# ------------------------------------------------------------------------------------------------------- #

# How to contribure t othe Project :

prerequisites : On a WSL instal golang and also git

Send a mail containing a motivation letter and github username to my mail adress, after a probably short 
period of reflection whe could add you to the list of colaborators of the project.
After that you could "git copy" the project.

And TaDam you can contribute to the project !

# ------------------------------------------------------------------------------------------------------- #

# Instruction for instalation and use :

prerequisites : On a WSL instal golang

On the github page of the project download it at .ZIP, after the download unzip the project in your WSL.

And TaDam tipe the comand "go run serveur.go" and go to localhost:8000 page and read the wiki on the site
to now how to use it.

# ------------------------------------------------------------------------------------------------------- #

# Technologies and information about prodject :

The project is in golang language and use :

 - math/rand package to generate randomly information abourt monster.
 - html/template and net/http to manage the web site.
 - encoding/json and io/ioutil to manage the json monster catalogue
 - strconv to make information about monster in string format need to have random number in it.

All monster generated are stocked in the Data/monster.json .

# ------------------------------------------------------------------------------------------------------- #

# FAQ about the project :

 - nothing

# ------------------------------------------------------------------------------------------------------- #

# Copyright and License Information

The project is developed for a personnal purpose, also the Copyright and License of the 5em edition dosen't
aply on the project. But if the project still grow we could do thomthing whith the SRD, OGL or Creative Commons
of Wizard of coast (DnD 5em) to publish a oficial website or something else...

(see https://dnd.wizards.com/fr/resources/systems-reference-document for more information about
SRD, OGL or Creative Commons)

# ------------------------------------------------------------------------------------------------------- #

# Actual contributor of the project :
 - Kelyan DANIS
 -
 -