package prome

type ATSMetrics struct {
	Global GlobalMetrics `json:"global"`
}

type GlobalMetrics struct {
	// QPS 相关指标
	IncomingRequests int64 `json:"proxy.process.http.incoming_requests" desc:"Total incoming requests per second"` // 总请求数
	GetRequests      int64 `json:"proxy.process.http.get_requests" desc:"Total GET requests"`                      // GET 请求数
	HeadRequests     int64 `json:"proxy.process.http.head_requests" desc:"Total HEAD requests"`                    // HEAD 请求数
	TraceRequests    int64 `json:"proxy.process.http.trace_requests" desc:"Total TRACE requests"`                  // TRACE 请求数
	OptionRequests   int64 `json:"proxy.process.http.options_requests" desc:"Total OPTIONS requests"`              // OPTIONS 请求数
	PostRequests     int64 `json:"proxy.process.http.post_requests" desc:"Total POST requests"`                    // POST 请求数
	PutRequests      int64 `json:"proxy.process.http.put_requests" desc:"Total PUT requests"`                      // PUT 请求数
	PushRequests     int64 `json:"proxy.process.http.push_requests" desc:"Total PUSH requests"`                    // PUSH 请求数
	DeleteRequests   int64 `json:"proxy.process.http.delete_requests" desc:"Total DELETE requests"`                // DELETE 请求数
	PurgeRequests    int64 `json:"proxy.process.http.purge_requests" desc:"Total PURGE requests"`                  // PURGE 请求数
	ConnectRequests  int64 `json:"proxy.process.http.connect_requests" desc:"Total CONNECT requests"`              // CONNECT 请求数

	// 并发相关指标
	CurrentClientConnections int64 `json:"proxy.process.http.current_client_connections" desc:"Current client connections"` // 当前客户端连接数

	// 吞吐相关指标
	UserAgentTotalResponseBytes    int64 `json:"proxy.process.http.user_agent_total_response_bytes" desc:"Total throughput in bytes per second"`                       // 用户代理总响应字节数
	OriginServerTotalResponseBytes int64 `json:"proxy.process.http.origin_server_total_response_bytes" desc:"Total throughput of original server in bytes per second"` // 源服务器总响应字节数

	// 响应延迟相关指标
	TotalTransactionsTime int64 `json:"proxy.process.http.total_transactions_time" desc:"Total transaction time in microseconds"` // 总事务处理时间（微秒）
	CompletedRequests     int64 `json:"proxy.process.http.completed_requests" desc:"Total completed requests"`                    // 完成的请求数

	// 数据大小分布相关指标
	ResponseDocumentSize100 int64 `json:"proxy.process.http.response_document_size_100" desc:"Number of responses with size 0-100 bytes"` // 响应数据大小在 0-100 字节的请求数
	ResponseDocumentSize1K  int64 `json:"proxy.process.http.response_document_size_1K" desc:"Number of responses with size 100B-1KB"`     // 响应数据大小在 100B-1KB 的请求数
	ResponseDocumentSize3K  int64 `json:"proxy.process.http.response_document_size_3K" desc:"Number of responses with size 1KB-3KB"`      // 响应数据大小在 1KB-3KB 的请求数
	ResponseDocumentSize5K  int64 `json:"proxy.process.http.response_document_size_5K" desc:"Number of responses with size 3KB-5KB"`      // 响应数据大小在 3KB-5KB 的请求数
	ResponseDocumentSize10K int64 `json:"proxy.process.http.response_document_size_10K" desc:"Number of responses with size 5KB-10KB"`    // 响应数据大小在 5KB-10KB 的请求数
	ResponseDocumentSize1M  int64 `json:"proxy.process.http.response_document_size_1M" desc:"Number of responses with size 10KB-1MB"`     // 响应数据大小在 10KB-1MB 的请求数
	ResponseDocumentSizeInf int64 `json:"proxy.process.http.response_document_size_inf" desc:"Number of responses with size 1M or more"`  // 响应数据大小在大于 1MB 的请求数

	// 缓存命中率相关指标
	CacheTotalHits        int64   `json:"proxy.process.cache_total_hits" desc:"Total bytes from cache hits"`           // 缓存总命中数
	CacheTotalMisses      int64   `json:"proxy.process.cache_total_misses" desc:"Total bytes from cache misses"`       // 缓存总未命中数
	CacheTotalHitsBytes   float64 `json:"proxy.process.cache_total_hits_bytes" desc:"Total bytes from cache hits"`     // 缓存命中总字节数
	CacheTotalMissesBytes float64 `json:"proxy.process.cache_total_misses_bytes" desc:"Total bytes from cache misses"` // 缓存未命中总字节数

	// 换成操作相关指标
	CacheOpLookups int64 `json:"proxy.process.http.cache_lookups" desc:"Number of cache lookups"` // 代理服务器进行的缓存查找次数
	CacheOpWrites  int64 `json:"proxy.process.http.cache_writes" desc:"Number of cache writes"`   // 代理服务器向缓存中写入新资源的次数
	CacheOpUpdates int64 `json:"proxy.process.http.cache_updates" desc:"Number of cache updates"` // 缓存中已有资源进行更新的次数
	CacheOpDeletes int64 `json:"proxy.process.http.cache_deletes" desc:"Number of cache deletes"` // 缓存中删除资源的次数

	// RAM 命中率相关指标
	RamCacheHits   float64 `json:"proxy.process.cache.ram_cache.hits" desc:"Total RAM cache hits"`     // RAM 缓存命中数
	RamCacheMisses float64 `json:"proxy.process.cache.ram_cache.misses" desc:"Total RAM cache misses"` // RAM 缓存未命中数

	// 资源占用相关指标
	CacheBytesUsed  int64 `json:"proxy.process.cache.bytes_used" desc:"Total bytes used in cache"`                     // 缓存已使用字节数
	CacheBytesTotal int64 `json:"proxy.process.cache.bytes_total" desc:"Total bytes available in cache"`               // 缓存总字节数
	RamCacheUsed    int64 `json:"proxy.process.cache.ram_cache.bytes_used" desc:"Total bytes used in RAM cache"`       // RAM 缓存已使用字节数
	RamCacheTotal   int64 `json:"proxy.process.cache.ram_cache.total_bytes" desc:"Total bytes available in RAM cache"` // RAM 缓存总字节数

	// 缓存性能相关指标
	KBReadPerSec  float64 `json:"proxy.process.cache.KB_read_per_sec" desc:"Cache read rate in KB per second"`   // 缓存每秒读取速率（KB）
	KBWritePerSec float64 `json:"proxy.process.cache.KB_write_per_sec" desc:"Cache write rate in KB per second"` // 缓存每秒写入速率（KB）
}
