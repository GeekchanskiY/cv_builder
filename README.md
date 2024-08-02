# cv builder


## ZEN of CV Builder 
Any microservice project may be monolith, but not all the monolith apps may be microservices \
More data is predefined is better, than less \
Sometimes even the same responsibility can conflict with itself \
The more relations present in the database - the better cv you will have in result \
Abuse the system, hate HR's, be a punk! \
If skill has no domains - it can be used everywhere \
If skill has a domain - it cant be used in others \
But it still can be used in a child domains \
You can regenerate current cv, but if you want something completely new - you should re-create it \

## HowTo

### Store the data
Import (and export in future) allows to share data in json packs, where all the data will be stored
and divided by topics. For example, you could have a system design pack, or go developer pack, AWS pack,
which will contain only the data about the topic you need. It's useful for manual writing skills, responsibilities,
and so on, but under the hood all this data will be stored in one place, and will not be divided into groups at all.

### Fill the data

#### Responsibilities 
id skill_id, name, priority, comments, experience

Priority: 1-10 mark of how often you should see this responsibility, where 1 - never, 10 - almost always

Experience: experience years to achieve 'mastery' of this feature, or which age of developer could implement
this feature more often. For example, it's 1 for writing business logic, but 7 for system design.

skill_id/skill_name: reference to skill, which this responsibility belongs to.

comments: they are used for frontend only. Fill them as you wish just to make it easier to prepare for the interview.

name: the name which will probably displayed in the CV, if there are no synonyms selected for it.
