module mod_file

go 1.22.2

// mod file are really important, beacuse of these three reasons
// first, it holds the project identification, which means it defins the project name along with the code hosting provider as the prefix 
// second, it has a version control, which tells about the golang version being used thus maintaining consistency across different envs
// third, it acts likes a registry talking about all the external packages which the project requires, thus other devs can easily install and setup the dependencies 