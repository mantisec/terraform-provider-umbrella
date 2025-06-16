output "admin_user_info" {
  description = "Information about the admin user"
  value = {
    id                   = umbrella_users.admin_user.id
    user_id              = umbrella_users.admin_user.user_id
    email                = umbrella_users.admin_user.email
    firstname            = umbrella_users.admin_user.firstname
    lastname             = umbrella_users.admin_user.lastname
    role                 = umbrella_users.admin_user.role
    role_id              = umbrella_users.admin_user.role_id
    timezone             = umbrella_users.admin_user.timezone
    status               = umbrella_users.admin_user.status
    two_factor_enabled   = umbrella_users.admin_user.two_factor_enabled
    last_login_time      = umbrella_users.admin_user.last_login_time
  }
}

output "readonly_user_info" {
  description = "Information about the read-only user"
  value = {
    id                   = umbrella_users.readonly_user.id
    user_id              = umbrella_users.readonly_user.user_id
    email                = umbrella_users.readonly_user.email
    firstname            = umbrella_users.readonly_user.firstname
    lastname             = umbrella_users.readonly_user.lastname
    role                 = umbrella_users.readonly_user.role
    role_id              = umbrella_users.readonly_user.role_id
    timezone             = umbrella_users.readonly_user.timezone
    status               = umbrella_users.readonly_user.status
    two_factor_enabled   = umbrella_users.readonly_user.two_factor_enabled
    last_login_time      = umbrella_users.readonly_user.last_login_time
  }
}

output "reporting_user_info" {
  description = "Information about the reporting user"
  value = {
    id                   = umbrella_users.reporting_user.id
    user_id              = umbrella_users.reporting_user.user_id
    email                = umbrella_users.reporting_user.email
    firstname            = umbrella_users.reporting_user.firstname
    lastname             = umbrella_users.reporting_user.lastname
    role                 = umbrella_users.reporting_user.role
    role_id              = umbrella_users.reporting_user.role_id
    timezone             = umbrella_users.reporting_user.timezone
    status               = umbrella_users.reporting_user.status
    two_factor_enabled   = umbrella_users.reporting_user.two_factor_enabled
    last_login_time      = umbrella_users.reporting_user.last_login_time
  }
}

output "bypass_user_info" {
  description = "Information about the bypass user"
  value = {
    id                   = umbrella_users.bypass_user.id
    user_id              = umbrella_users.bypass_user.user_id
    email                = umbrella_users.bypass_user.email
    firstname            = umbrella_users.bypass_user.firstname
    lastname             = umbrella_users.bypass_user.lastname
    role                 = umbrella_users.bypass_user.role
    role_id              = umbrella_users.bypass_user.role_id
    timezone             = umbrella_users.bypass_user.timezone
    status               = umbrella_users.bypass_user.status
    two_factor_enabled   = umbrella_users.bypass_user.two_factor_enabled
    last_login_time      = umbrella_users.bypass_user.last_login_time
  }
}

output "international_user_info" {
  description = "Information about the international user"
  value = {
    id                   = umbrella_users.international_user.id
    user_id              = umbrella_users.international_user.user_id
    email                = umbrella_users.international_user.email
    firstname            = umbrella_users.international_user.firstname
    lastname             = umbrella_users.international_user.lastname
    role                 = umbrella_users.international_user.role
    role_id              = umbrella_users.international_user.role_id
    timezone             = umbrella_users.international_user.timezone
    status               = umbrella_users.international_user.status
    two_factor_enabled   = umbrella_users.international_user.two_factor_enabled
    last_login_time      = umbrella_users.international_user.last_login_time
  }
}

output "all_users_summary" {
  description = "Summary of all created users"
  value = {
    admin_user = {
      email = umbrella_users.admin_user.email
      role  = umbrella_users.admin_user.role
    }
    readonly_user = {
      email = umbrella_users.readonly_user.email
      role  = umbrella_users.readonly_user.role
    }
    reporting_user = {
      email = umbrella_users.reporting_user.email
      role  = umbrella_users.reporting_user.role
    }
    bypass_user = {
      email = umbrella_users.bypass_user.email
      role  = umbrella_users.bypass_user.role
    }
    international_user = {
      email = umbrella_users.international_user.email
      role  = umbrella_users.international_user.role
    }
  }
}

output "user_count" {
  description = "Total number of users created"
  value       = 5
}

output "users_by_role" {
  description = "Users grouped by role"
  value = {
    full_admin = [
      {
        email = umbrella_users.admin_user.email
        name  = "${umbrella_users.admin_user.firstname} ${umbrella_users.admin_user.lastname}"
      }
    ]
    read_only = [
      {
        email = umbrella_users.readonly_user.email
        name  = "${umbrella_users.readonly_user.firstname} ${umbrella_users.readonly_user.lastname}"
      },
      {
        email = umbrella_users.international_user.email
        name  = "${umbrella_users.international_user.firstname} ${umbrella_users.international_user.lastname}"
      }
    ]
    block_page_bypass = [
      {
        email = umbrella_users.bypass_user.email
        name  = "${umbrella_users.bypass_user.firstname} ${umbrella_users.bypass_user.lastname}"
      }
    ]
    reporting_only = [
      {
        email = umbrella_users.reporting_user.email
        name  = "${umbrella_users.reporting_user.firstname} ${umbrella_users.reporting_user.lastname}"
      }
    ]
  }
}