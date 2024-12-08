# giftexchanger
go script to randomly assign the family gift exchange for kids respecting pairs and 1st cousins

# Gift Exchange Script
This Go script creates a gift exchange mapping for a list of kids' names, ensuring that no name maps to itself and no pairs are made that would violate the specified invalid pairs.  

## Features
- Shuffles a list of names to create a gift exchange mapping.
- Ensures that no name maps to itself.
- Ensures that no pairs are made that would violate the specified invalid pairs.

## Usage
- Install Go: Ensure you have Go installed on your machine. You can download it from golang.org.  
- Clone the Repository: Clone this repository to your local machine.  
- Modify the Names List: Update the names slice in the main function with the list of kids' names.  
- Modify Invalid Pairs: Update the invalidPairs map with any pairs that should not be allowed.  
- Run the Script: Execute the script using the following command:  
```bash
go run main.go
```
