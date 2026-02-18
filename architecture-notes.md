1. Routes
- path = root/api/routes/<name>.go
- create new routes
<!-- - import = 
fiber router "github.com/gofiber/fiber/v2"
app "github.com/username-owner-repo/api/app" ->modules in the projects
service "github.com/username-owner-repo/pkg/<name>/app"
handlers "github.com/username-owner-repo/api/handlers/<name>" -->

- create routes function <Name>Router
    inside the routes, there are 5 endpoints (create, read, read detail, update, delete)
- Registering the new routes (function) to routes/routes.go

2. Model
- path : pkg/entities/<name>.go
- structure of the tables

3. Handlers
- path = root/api/handlers/<name>/
    generate 5 handler (1 for each endpoints in routes)
    create_<name>_handlers.go, etc
- pattern : will be there diff if it has to read the body request
<!-- - import :  -->
3.1 Read (dont think about the filter and paginate first)

4. Service
- path = pkg/<name>/service/
- ignore unit test first
- generate 5 services (1 for each endpoints in routes)
    create_<name>_service.go, etc
    + service.go

- service.go : 
    consist of interface, struct, func NewService
    register all 5 service to the interface
- create_<name>_service.go, (update, read, read_detail, delete)




========== Pak Wahyu ==========
bisa handle table ber relasi
handling date di react (datetime input)
filter
all columns filterable, tinggal menyesuaikan kebutuhan FSD/user
tanggal filter : from-to
tipe input menyesuaikan tipe data
input table berupa migration
================================