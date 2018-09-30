

start as wet
exit 1 when wet
exit 0 when dry

( wet -> dry ) : airdry 
( dry -> wet ) : rain 
(wet -> frozen) : freeze
(frozen -> wet) : met


