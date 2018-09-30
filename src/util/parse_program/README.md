
- Grammar 
 
( wet -> dry ) : airdry              --> rule --> statement
( dry -> wet ) : rain 	
(wet -> frozen) : freeze
(frozen -> wet) : met

start as wet						--> start --> statement
exit 1 when wet						
exit 0 when dry							--> exit
when - dry the wet : someexternalevent	--> hook --> statement


grammar:

string: a-z, a-z
axiom: 1+ :string
axiom: 1+ :axiom

rule: (:axiom->:axiom) : :axiom
rule: (:axiom->:axiom) - :hookaxiom : :axiom
start: start as :axiom
code: [0,1]
exit: exit :code when :axiom
script: :axiom
hookaxiom: axiom
hook: when - :hookaxiom : :script

statement: :script
statement: :exit
statement: :hook
statement: :rule

program: 1+ :statement