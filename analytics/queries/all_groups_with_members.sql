select gm.Name,
    gm.Type,
    gm.Location,
    gm.MemberLocation,
    case when u.id is not null then 'User'
        when g.id is not null then 'Group'
        else 'Unknown'
    end as MemberType,
    case when u.name is not null then u.name
        when g.name is not null then g.name
        else null
    end as MemberName,
    case when u.givenname = u.surname then trim(u.givenname)
        when u.givenname is null and u.name is not null then trim(u.name)
        when u.givenname is not null then trim(u.givenname) || ' ' || trim(u.surname)
        when u.name is not null then u.name
        when g.name is not null then g.name
        else null
    end as MemberFullName
from(
    select id as Id, 
        name as Name,
        location as Location,
        type as Type,
        member as MemberLocation
    from identity_intelligence_prd."group"
CROSS JOIN UNNEST(members) AS t(member)
) as gm
left join identity_intelligence_prd."user" as u on gm.MemberLocation = u.location
left join identity_intelligence_prd."group" as g on gm.Id = g.id