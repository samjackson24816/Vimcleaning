# Vimcleaning



The ultimate casual game to play 15 seconds at a time while waiting for an agent to finish


Made in Go

Will be a series of randomly generated grid-based rooms where you are a character and you move around with vim keybinds and pick up and place characters to organize the room, picking up all the letters of a word to form the word.

You can't pick up a letter and then pick up another one, and you can't copy multiple letters
You should be able to use b to jump back to the next left empty square (like vim) and w to jump forward to the right before the next right empty square

Words are generated from a database of commonly used words

The word generation should be instant---the idea is you bring up the window with a hotkey and then get rid of it the moment you have something else to do.  The whole game should be centered around that idea

I want to make this with no agents.  Learn by hand



Ideas for new features:

Timer that times your progess in letters per min per round

Replay the same word with the same seed or continue with different keybinds

Show previous time and stats in hud on top or bottom

Add blockers that you can't move over

Rate the number of moves a user takes and compare to optimal path generated with A* (maybe on pregenerated maps)

Use seeds to genereate deterministic run and record user times against seeds?
