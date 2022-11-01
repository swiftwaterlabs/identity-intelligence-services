select u.id,
    u.directory,
    u.name,
    u.location,
    u.type,
    u.objecttype,
    u.upn,
    u.givenname,
    u.surname,
    trim(u.givenname) || ' ' || trim(u.surname) as fullname,
    u.email,
    u.company,
    u.department,
    u.title,
    case 
        when u.type = 'SAM_MACHINE_ACCOUNT' then 'Machine Account'
        when u.location like '%OU=Service Account%' then 'Service Account'
        when u.location like '%OU=Resource Account%' then 'Resource Account'
        when u.location like '%OU=Admin Account%' then 'Admin Account'
        else 'User Account'
    end as user_type,
    case 
        when u.type = 'SAM_MACHINE_ACCOUNT' then true
        when u.location like '%OU=Service Account%' then true
        when u.location like '%OU=Resource Account%' then true
        when u.location like '%OU=Admin Account%' then false
        else false
    end as is_non_human,
    case
         when u.location like '%OU=Terminated%' then 'Terminated'
         when u.location like '%OU=Active%' then 'Active'
         when u.location like '%OU=Non_Active%' or u.location like '%OU=Disabled%' then 'Disabled'
         when u.type = 'SAM_MACHINE_ACCOUNT' or u.location like '%OU=Service Account%' or u.location like '%OU=Resource Account%' then 'Active'
         else 'Unknown'
    end as user_status,
    trim(u.manager)!='' and u.manager is not null as has_manager,
    trim(um.location)!='' and um.location is not null as has_valid_manager,
    u.manager as manager_raw,
    um.location as manager_location,
    um.name as mamanger_name,
    um.upn as manager_upn,
    trim(um.givenname) || ' ' || trim(um.surname) as manager_fullname,
    case 
        when um.type = 'SAM_MACHINE_ACCOUNT' then 'Machine Account'
        when um.location like '%OU=Service Account%' then 'Service Account'
        when um.location like '%OU=Resource Account%' then 'Resource Account'
        when um.location like '%OU=Admin Account%' then 'Admin Account'
        else 'User Account'
    end as manager_user_type,
    case
         when um.location like '%OU=Terminated%' then 'Terminated'
         when um.location like '%OU=Active%' then 'Active'
         when um.location like '%OU=Non_Active%' or um.location like '%OU=Disabled%' then 'Disabled'
         when um.type = 'SAM_MACHINE_ACCOUNT' or um.location like '%OU=Service Account%' or um.location like '%OU=Resource Account%' then 'Active'
         else 'Unknown'
    end as manager_status
from user u
left join user um on u.manager = um.location
