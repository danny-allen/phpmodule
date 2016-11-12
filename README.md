# PHPMODULE

<br />


### Creating a module - IN PROGRESS

Format: `program action name`


```{r, engine='sh', count_lines}
$ phpmodule init Billing    # Create Module
```

<br />

---

<br />


### Creating module components - NOT IMPLEMENTED YET

Format: `program action type (name)`

	
```{r, engine='sh', count_lines}
# Create Module Data
$ phpmodule create data Customer

# Create Module Enum
$ phpmodule create enum PlanCode

# Create Module Exception
$ phpmodule create exception (name) 	# Name Optional

# Create Module Test
$ phpmodule create test Customer
```