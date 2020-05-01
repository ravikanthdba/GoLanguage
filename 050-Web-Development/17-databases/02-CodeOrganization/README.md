3 approaches:

Approach 1:

code_dir:
    templates
    main.go
    
Approach 2:

code_dir:
    models
        datastructure
        operations
    templates
        *.gohtml
    main.go
    
Approach 3:
   
code_dir
    models
        handlers.go - Controllers
        models.go - datastructure and operations on the DS
    templates
        *.gohtml - WEB files
    config (drivers)
        db.go
        template.go (initializing the template)
    main.go
    
    