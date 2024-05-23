import http from '@/plugins/axios'


//推送一条示例数据
export const postDemoData = (namespace, data) => http.Post("/go/" + namespace, data)

//劫持过境数据
export const queryJson = (namespace, params) => http.Get("/api/v1/" + namespace + "/json", params)


//用户变量相关（完成数据校验）
export const queryVars = (namespace, params) => http.Get("/api/v1/" + namespace + "/vars", params) //查询用户变量映射
export const postVars = (namespace, data) => http.Post("/api/v1/" + namespace + "/vars", data) //添加用户变量映射
export const queryFlattening = (namespace, params) => http.Get("/api/v1/" + namespace + "/flattening", params) //展开过境数据

//用户模板相关
export const getTemplate = (namespace, params) => http.Get("/api/v1/" + namespace + "/template", params)
export const postTemplate = (namespace, data) => http.Post("/api/v1/" + namespace + "/template", data)


//客户端相关
export const getClient = (namespace, params) => http.Get("/api/v1/" + namespace + "/client", params)
export const postClient = (namespace, data) => http.Post("/api/v1/" + namespace + "/client", data)
export const getClientOne = (namespace, id, params) => http.Get("/api/v1/" + namespace + "/client/" + id, params)
export const putClientOne = (namespace, id, data) => http.Put("/api/v1/" + namespace + "/client/" + id, data)
export const putClientInfoOne = (namespace, id, data) => http.Put("/api/v1/" + namespace + "/client-info/" + id, data)
export const deleteClientOne = (namespace, id, params) => http.Delete("/api/v1/" + namespace + "/client/" + id, params)


//命名空间（完成数据校验）
export const getNamespace = (params) => http.Get("/api/v1/namespace", params)
export const postNamespace = (data) => http.Post("/api/v1/namespace", data)
export const getNamespaceOne = (id, params) => http.Get("/api/v1/namespace/" + id, params)
export const deleteNamespaceOne = (id) => http.Delete("/api/v1/namespace/" + id)
export const putNamespaceOne = (id, data) => http.Put("/api/v1/namespace/" + id, data)


//定时消息相关（完成数据校验）
export const getCrontab = (params) => http.Get("/api/v1/crontab", params)
export const postCrontab = (data) => http.Post("/api/v1/crontab", data)
export const getCrontabOne = (id, params) => http.Get("/api/v1/crontab/" + id, params)
export const deleteCrontabOne = (id) => http.Delete("/api/v1/crontab/" + id)
export const putCrontabOne = (id, data) => http.Put("/api/v1/crontab/" + id, data)


//用户登录相关
export const login = (data) => http.Post("/auth/login", data)
export const logout = (data) => http.Post("/auth/logout", data)
export const updatePassword = (id, data) => http.Put("/api/v1/auth/user/password/" + id, data)
