#任务
描述一个向导服务，服务名为RouteGuide
定义四种不同的信息类型，分别为Point,Rectangle,Feature,RouteSummary以及Chat。
定义四个方法：
- GetFeature(输入为一个Point,返回这个点的Feature)
- ListFeatures(输入为一个Rectangle,输出流这个区域内所有的Feature)
- RecordRoute(输入流为每个时间点的位置Point,返回一个RouteSummary)
- Recommend(输入流RecommendationRequest,返回一个输出流Feature)