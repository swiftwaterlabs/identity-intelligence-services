select name as Name,
    location as Location,
    type as Type,
    member as MemberLocation
from "group"
CROSS JOIN UNNEST(members) AS t(member)