

start as wet
exit 1 when wet
exit 0 when dry
when - dry the wet : someexternalevent
( wet -> dry ) : airdry 
( dry -> wet ) : rain 
(wet -> frozen) : freeze
(frozen -> wet) : met
