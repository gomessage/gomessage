import http from '@/plugins/axios'

const namespace = "default"


//推送一条demo数据
export const postDemoData = data => http.Post("/go/" + namespace, data)

//劫持过境数据
export const queryJson = params => http.Get("/api/v1/" + namespace + "/json", params)

//用户变量相关
export const queryVars = params => http.Get("/api/v1/" + namespace + "/vars", params)
export const postVars = data => http.Post("/api/v1/" + namespace + "/vars", data)

//用户模板相关
export const getTemplate = params => http.Get("/api/v1/" + namespace + "/template", params)
export const postTemplate = data => http.Post("/api/v1/" + namespace + "/template", data)


//客户端相关
export const getClient = params => http.Get("/api/v1/" + namespace + "/client", params)
export const postClient = data => http.Post("/api/v1/" + namespace + "/client", data)
export const getClientOne = (id, params) => http.Get("/api/v1/" + namespace + "/client/" + id, params)
export const putClientOne = (id, data) => http.Put("/api/v1/" + namespace + "/client/" + id, data)
export const deleteClientOne = (id, params) => http.Delete("/api/v1/" + namespace + "/client/" + id, params)

