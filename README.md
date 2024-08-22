# forum-system 后端代码
本项目基于Gin框架开发，是一个以投票功能为核心的web论坛项目，实现了用户注册、登陆、发帖、社区分类、获取帖子详情以及用户投票等功能，并实现了基于用户投票的帖子排名功能。
项目在结构上分为3层，首先是server层，这是服务的入口，负责处理路由、参数校验和请求转发；然后是service层，这是服务层，负责处理业务逻辑；最后是存储层，负责数据和存储相关的功能，主要是操作mysql和redis。项目实现了用户注册、登陆、发帖、社区分类、获取帖子详情以及用户投票等功能，并实现了基于用户投票的帖子排名功能。在登陆验证方面，使用了JWT，集成Access Token和Refresh Token机制，确保令牌的有效性和安全性；在生成用户id与帖子id方面，使用雪花算法解决了分布式系统中的ID生成瓶颈问题, 提高了系统的性能和可扩展性；在项目核心功能投票和投票排名方面，使用Redis实现了帖子投票功能，使用Zset数据结构存储了发布时间，以及综合发布时间与投票情况的分数，实现了帖子按时间排名和按热度排名，并抑制了排名的“马太效应”；在限流算法方面，选择了令牌桶限流的算法，实现了项目中用户请求速率的平滑控制与突发访问的合理接纳。
