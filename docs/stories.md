# Soullocke Tracker App User Stories

## Soullocke Concept

- Players start a run
- Players encounter a pokemon on a location
    - Outcome A: Everyone catches the pokemon => Successful link, can be used, location locked
    - Outcome B: One fails to catch the pokemon => Failed link, cant be used, location locked
- Players try to advance as much as possible
- Whenever a pokemon dies => the linked pokemon are marked as unavailable and cannot be used anymore
- If a team wipes => Run has failed, a new run will be started
- If the win condition is reached (usually top 4 beaten) => successful run

## Ungrouped

Als Spieler möchte ich eine Soullocke Lobby eröffnen, um Pokemon Encounter zu tracken.

Als Lobbyersteller möchte ich andere Spieler zu meiner Lobby einladen, um gemeinsam an einem Run zu kollaborieren.

Als Lobbyersteller möchte ich die Lobby vor fremden Zugriff schützen, damit nur berechtigte Personen Änderungen vornehmen können.

Als Spieler möchte ich jederzeit die Lobby öffnen können, um jederzeit den aktuellen Status einsehen zu können.

Als Spieler möchte ich eine Game Edition auswählen, in der ich spiele, um Locations und deren Encounter zur Verfügung zu haben.

Als Spieler möchte ich eine Location für ein Encounter auswählen, um sich zu merken, welche Locations ich bereits habe.

Als Spieler möchte ich das Pokemon für ein Encounter auswählen, um zu tracken, welches Pokemon an der Location das Encounter war.

Als Spieler möchte ich den Ausgang meines Encounters erfassen, damit das System die Verfügbarkeit des Pokémon und den Status der Location korrekt fortschreibt.

Als Spieler möchte ich den Tod eines Pokémon erfassen, damit alle davon abhängigen gelinkten Pokémon automatisch ebenfalls als nicht mehr nutzbar markiert werden.

Als Spieler möchte ich eintragen, wann ich ein Encounter nicht fangen konnte, damit ich eine Location als "locked" markieren kann.

Als Spieler möchte ich Encounter bearbeiten können, damit ich falsche Einträge korrigieren kann.

Als Spieler möchten wir unsere Encounter derselben Location zu einem gemeinsamen Link zusammenfassen, damit Abhängigkeiten zwischen unseren Pokémon automatisch nachvollziehbar sind.

Als Spieler möchte ich Spitznamen für die gefangenen Encounter vergeben, um die Nuzlocke Regeln einzuhalten.

Als Spieler wollen wir Pokemon als Teil unseres Teams markieren, um Team Management zu erhalten.

Als Spieler wollen wir Pokemon als Teil unserer Boxen markieren, um zu sehen, welche Backups wir verfügbar haben.

Als Spieler möchte ich Änderungen meiner Mitspieler in Echtzeit sehen, damit wir unseren Run synchron verfolgen können.

Als Spieler möchte ich sehen, welche Pokemon miteinander verlinkt sind, um direkt einsehen zu können, welche Pokemon zueinander gelinked sind.

Als Spieler möchte ich einen Run in einer Lobby anlegen, um den Fortschritt eines Runs zu verfolgen.

Als Spieler möchte ich einen Run als fehlgeschlagen markieren, wenn wir z. B. wipen, um den Run zu beenden.

Als Spieler möchte ich einen neuen Run anlegen, um wieder von vorne zu beginnen.

Als Spieler möchte ich einen Run abschließen, um den Erfolg der Lobby zu markieren.

Als Spieler möchte ich vergangene Runs einsehen können, um zu sehen, wie die Historie einer Lobby war.

Als Spieler möchten wir bereits locked Locations einsehen können, um doppelte Encounter zu vermeiden.

Als Spieler möchte ich noch offene Encounter erblicken können, um meine Encounter zu vervollständigen.

Als Spieler möchte ich den Grund für unbrauchbare Pokemon einsehen können, um die Historie nachzuvollziehen.

Als Spieler möchte ich neben Locations auch gifted Pokemon eintragen können, um diese gesondert von den statischen Encountern zu tracken. 

Als Spieler möchte ich sehen, welche Spieler für eine Location bereits einen Encounter erfasst haben und welche noch fehlen, damit wir Links vollständig abschließen können.

Als Spieler möchte ich sehen, welche Locations bereits verwendet oder gesperrt sind, damit keine doppelten Encounter entstehen.

Als Spieler möchte ich erkennen, ob ein Encounter noch unvollständig ist, damit offene Einträge nicht übersehen werden.



## Utility

Als Spieler möchte ich direkt eine Typen Matrix sehen, um mir eine schnelle Übersicht zu geben.

Als Spieler möchte ich die Level Caps einsehen können, um eine schnelle Übersicht zu haben. 