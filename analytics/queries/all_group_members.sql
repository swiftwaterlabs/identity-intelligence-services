select name as Name,
    location as Location,
    type as Type,
    member as MemberLocation
from identity_intelligence_prd."group"
CROSS JOIN UNNEST(members) AS t(member)