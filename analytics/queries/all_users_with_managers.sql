select u.id,
    u.directory as Directory,
    u.name as Name,
    u.location as Location,
    u.type as Type,
    u.objecttype as ObjectType,
    u.upn as Upn,
    u.givenname as GivenName,
    u.surname as SurName,
    trim(u.givenname) || ' ' || trim(u.surname) as FullName,
    u.email as Email,
    u.company as Company,
    u.department as Department,
    u.title as JobTitle,
    case 
        when u.type = 'SAM_MACHINE_ACCOUNT' then 'Machine Account'
        when u.location like '%OU=Service Account%' then 'Service Account'
        when u.location like '%OU=Admin Account%' then 'Admin Account'
        when u.location like '%OU=Resource Account%' then 'Resource Account'
        else 'User Account'
    end as UserType,
    case 
        when u.type = 'SAM_MACHINE_ACCOUNT' then true
        when u.location like '%OU=Service Account%' then true
        when u.location like '%OU=Admin Account%' then false
        when u.location like '%OU=Resource Account%' then true
        else false
    end as IsNonHumanAccount,
    case
         when u.location like '%OU=Terminated%' then 'Terminated'
         when u.location like '%OU=Active%' then 'Active'
         when u.location like '%OU=Non_Active%' or u.location like '%OU=Disabled%' then 'Disabled'
         when u.type = 'SAM_MACHINE_ACCOUNT' or u.location like '%OU=Service Account%' or u.location like '%OU=Resource Account%' then 'Active'
         else 'Unknown'
    end as UserStatus,
    trim(u.manager)!='' and u.manager is not null as HasManager,
    trim(um.location)!='' and um.location is not null as HasValidManager,
    u.manager as ManagerLocationRaw,
    um.location as ManagerLocation,
    um.name as ManagerName,
    um.upn as ManagerUpn,
    trim(um.givenname) || ' ' || trim(um.surname) as ManagerFullName,
    case 
        when um.type is null or um.type = '' then null
        when um.type = 'SAM_MACHINE_ACCOUNT' then 'Machine Account'
        when um.location like '%OU=Service Account%' then 'Service Account'
        when um.location like '%OU=Admin Account%' then 'Admin Account'
        when um.location like '%OU=Resource Account%' then 'Resource Account'
        else 'User Account'
    end as ManagerUserType,
    case
         when um.location is null or um.location = '' then 'Missing'
         when um.location like '%OU=Terminated%' then 'Terminated'
         when um.location like '%OU=Active%' then 'Active'
         when um.location like '%OU=Non_Active%' or um.location like '%OU=Disabled%' then 'Disabled'
         when um.type = 'SAM_MACHINE_ACCOUNT' or um.location like '%OU=Service Account%' or um.location like '%OU=Resource Account%' then 'Active'
         else 'Unknown'
    end as ManagerStatus
from user u
left join user um on u.manager = um.location
