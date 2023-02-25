import http from '@/plugins/axios'

const namespace = "default"


//推送一条demo数据
export const postDemoData = data => http.Post("/go/" + namespace, data)

//劫持过境数据
export const queryJson = params => http.Get("/api/v1/" + namespace + "/json", params)

//用户变量相关
export const queryVars = (namespace, params) => http.Get("/api/v1/" + namespace + "/vars", params)
export const postVars = (namespace, data) => http.Post("/api/v1/" + namespace + "/vars", data)

//用户模板相关
export const getTemplate = (namespace, params) => http.Get("/api/v1/" + namespace + "/template", params)
export const postTemplate = (namespace, data) => http.Post("/api/v1/" + namespace + "/template", data)


//客户端相关
export const getClient = (namespace, params) => http.Get("/api/v1/" + namespace + "/client", params)
export const postClient = (namespace, data) => http.Post("/api/v1/" + namespace + "/client", data)
export const getClientOne = (id, params) => http.Get("/api/v1/" + namespace + "/client/" + id, params)
export const putClientOne = (id, data) => http.Put("/api/v1/" + namespace + "/client/" + id, data)
export const deleteClientOne = (id, params) => http.Delete("/api/v1/" + namespace + "/client/" + id, params)


//命名空间相关
export const getNamespace = (params) => http.Get("/api/v1/namespace", params)
