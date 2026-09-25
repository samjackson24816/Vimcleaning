# Vimcleaning



The ultimate casual game to play 15 seconds at a time while waiting for an agent to finish


Made in Go

Will be a series of randomly generated grid-based rooms where you are a character and you move around with vim keybinds and pick up and place characters to organize the room, picking up all the letters of a word to form the word.

You use x o cut and hold a letter and p to paste the letter (just to match with vim)
You can't pick up a letter and then pick up another one, and you can't copy multiple letters
You should be able to use b to jump back to the next left empty square (like vim) and w to jump forward to the right before the next right empty square

When you do escape-:-w-enter the room should save if it is done and a new room should be generated

Words are generated from a database of commonly used words

The word generation should be instant---the idea is you bring up the window with a hotkey and then get rid of it the moment you have something else to do.  The whole game should be centered around that idea


