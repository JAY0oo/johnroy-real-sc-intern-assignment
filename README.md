# SafetyCulture Internship Take Home Assessment

This is my solution for SafetyCulture's Take home Assessment challange.

## Component Goals

### Component 1: GetAllChildFolders
- Implement the `GetAllChildFolders` method to return all child folders of a given folder.
- Implement necessary error handling.
- Write unit tests for all methods in `get_folder.go`.

### Component 2: MoveFolder
- Implement the `MoveFolder` method to move a folder from one parent to another.
- Implement necessary error handling.
- Write unit tests for the `MoveFolder` method in `move_folder_test.go`.

## Results
- My implementation for both the `GetAllChildFolders` and `MoveFolder` work as expected, passing all unit tests.
- However, I had to introduce an additional parameter, `orgID`, to the `MoveFolder` method. This was done to maintain consistency, as `orgID` is also required by functions in Component 1 (e.g., `GetAllChildFolders`). Additionally, I encountered difficulties accessing `defaultOrgID`, so having `orgID` as a parameter ensures that I can consistently access and work with the correct organization ID across both components.

  
### Sample output 1 - Folders of orgID
These output is the program using a smaller sample size for easier demonstration.
- This is the initial printout of all folders belonging to the organization ID.
- Notice how "golf" has no subfolders.
- ![image](https://github.com/user-attachments/assets/c5ff974b-cbc0-492b-a179-be6ac7245165)
  
### Sample output 2 - Child folders of "alpha"
- This output are the subfolders of folder "alpha".
- ![image](https://github.com/user-attachments/assets/216ab5fb-516f-4914-8b8f-f24369e1b54d)

### Sample output 3 - Moving "alpha" to "golf"
- This output moves the folder "alpha" and all its sub folder to "golf".
- After the move, it shows the new folder structure.
- Note: The order of the folders might not be the same as before the move, but this doesn’t affect the functionality since the paths and hierarchy are correct.
- ![image](https://github.com/user-attachments/assets/db4d0afa-fecc-4020-96e0-3b49e05ccfe2)


### Sample output 4 - List "golf" subfolder
- Using `GetAllChildFolders` list all "golf" new subfolders.
- ![image](https://github.com/user-attachments/assets/f3c57e51-23d1-41d1-ab0d-e18d621e705d)

## Conclusion

Thank you for the opportunity, this was a great learning experience which pushed me to learn Golang. I'm grateful for the chance to showcase my work and for the challenges that helped me grow as a developer.

Best regards,  
Johnroy Real | [LinkedIn](www.linkedin.com/in/johnroy-real).



