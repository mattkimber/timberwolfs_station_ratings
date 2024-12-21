Timberwolf's Station Ratings
----------------------------

This is a small addon to change station rating calculations
in a way which better suits Timberwolf's Trains and Road
Vehicles, especially in games with an early start year.

Also due to the way in which cargo overrides operate, it
supplies 2x cargo sprites for all FIRS/AIRS cargo, so cargo 
does not appear as a generic box.

How it works
------------

As early vehicles are slow and cargo wagons are generally
slower than the default vehicles, the speed of vehicle
required to get a full ratings boost will increase over time.

Passengers: Starts at 30km/h (19mph) in 1830, rises to 
            200km/h (125mph) by 2000.
            
Food/mail: Starts at 20km/h (12mph) in 1830, rises to
           130km/h (81mph) by 2000.

Other cargo: Starts at 20km/h (12mph) in 1830, rises to
             112km/h (70mph) by 2000.

Prior to 1830 the minimum values are used, and after 2000 the
maximum values are used.

The rating is also adjusted to account for the low capacities
and likelihood of infrequent service early in the game.

Prior to 1900 the number of days since last service is evaluated
more leniently, and larger amounts of waiting cargo are allowed 
before a ratings penalty. This effect is increased for dates
before 1850; people were clearly impressed just to see a train!

This is balanced by stricter treatment of waiting cargo from 1970
onward, with increased strictness from 2010. Cargo stations will
struggle to achieve maximum rating unless you always have a train
waiting to load cargo. Passenger stations may have a small number
of passengers waiting to allow interchange, but will suffer heavy
penalties for overcrowding.

As the default age mechanics make little sense with vehicles having
20+ year lifespans, there is a small extra bonus for exceptional 
service compared to the default rating algorithm, to negate the 
need to constantly replace vehicles and/or advertise to get full 
station ratings at well-served stations.

If you are using a NewGRF industry set, this set will need to be
loaded after it for custom ratings to work. This will be detected
for several common industry sets.

Penalty for boring networks
---------------------------

A common frustration of OpenTTD games is that a small number of air 
and/or sea routes can easily outperform a complex rail and road
network financially, due to high cargo payments at 1/1 planespeed and
the effectively infinite capacity of ship docks.

This can lead to less interesting games overall, where high incomes
can be gained despite building very little infrastructure.

The penalty for boring networks attempts to compensate for this by
applying a ratings penalty to docks and airports whenever the last
vehicle to load cargo was a ship or aircraft. This means that players
must serve their docks and airports with other transport methods to
make use of their full capacity - either by having multiple vehicle
types loading at the same station and managing ships to not wait
extended periods of time, or by transporting the first leg via rail,
road or tram and only using ships as the intermediate stage of a
journey.

For supported industry sets, cargo which is obtained from offshore
industries is not affected by this penalty, allowing ships to be
used to bring it to shore.

The aircraft penalty is not applied to passengers or mail, as it
is usually not easy to build large airports close to a town, and 
connecting services will already need to be provided.

Note: as this is a somewhat unexpected use of NML, support for
industry sets has to be hardcoded on a set-by-set basis. Currently
only FIRS 4/5, AIRS, and the base game are supported.

Compatibility may be broken if industry sets introduce new cargo
types, or otherwise reorder their cargo list.

Reasonable requests to add additional industry sets will be
accepted, providing source code is available.