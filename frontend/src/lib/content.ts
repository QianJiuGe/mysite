export type BlogPost = {
  slug: string
  title: string
  summary: string
  tags: string[]
  publishedAt: string
  updatedAt: string
  readingMinutes: number
  sections: { heading: string; body: string }[]
}

export type ProjectCase = {
  id: string
  name: string
  summary: string
  tech: string[]
  role: string
  status: string
  challenge: string
  solution: string
  impact: string
  links: { label: string; href: string }[]
}

export const blogPosts: BlogPost[] = [
  {
    slug: 'ship-fast-with-guardrails',
    title: '在高迭代节奏下建立可回滚的交付护栏',
    summary: '用轻量规范、发布分层和可观测信号，把“快”从冒险变成可控能力。',
    tags: ['Engineering', 'Delivery'],
    publishedAt: '2026-03-18',
    updatedAt: '2026-04-02',
    readingMinutes: 8,
    sections: [
      {
        heading: '背景与问题',
        body: '团队在快速上线时最容易忽略回滚路径，导致线上风险被放大。只要把变更拆成可验证的小步，失败成本就会显著下降。',
      },
      {
        heading: '实践方案',
        body: '我把发布分为“预检、灰度、全量”三段，每段都有明确的观测指标和退出条件；同时把关键规则沉淀成模板，确保团队执行一致。',
      },
      {
        heading: '结果与复盘',
        body: '上线失败率下降后，团队对迭代节奏更有信心。后续重点是把更多回滚演练自动化，减少人为判断负担。',
      },
    ],
  },
  {
    slug: 'frontend-information-architecture',
    title: '个人站前端信息架构：从页面到演进路径',
    summary: '不是先画 UI，而是先确定信息层级、页面职责和可增长边界。',
    tags: ['Frontend', 'Architecture'],
    publishedAt: '2026-02-27',
    updatedAt: '2026-02-27',
    readingMinutes: 6,
    sections: [
      {
        heading: '为什么先做 IA',
        body: '视觉可以重做，但信息结构会影响长期维护成本。先定义页面职责，可以避免后期组件不断膨胀。',
      },
      {
        heading: '我的拆分方法',
        body: '将“首页、列表、详情”分成三类模板，每类只关注一种核心任务，再通过统一导航串联，保证迁移成本可控。',
      },
      {
        heading: '长期收益',
        body: '需求增长时，页面可以新增区块而不破坏现有结构，团队协作也更容易并行推进。',
      },
    ],
  },
  {
    slug: 'observability-notes-for-side-projects',
    title: '小型项目的可观测性笔记：先看见，再优化',
    summary: '日志、指标、告警不需要一步到位，但需要从第一天开始有最小闭环。',
    tags: ['SRE', 'Notes'],
    publishedAt: '2026-01-15',
    updatedAt: '2026-01-21',
    readingMinutes: 7,
    sections: [
      {
        heading: '最小闭环',
        body: '先有错误日志和基础指标，再扩展到延迟与吞吐。没有观察能力，优化只是猜测。',
      },
      {
        heading: '实践建议',
        body: '给每个关键接口定义成功率和耗时阈值，异常时第一时间定位到责任链路，减少故障排查时间。',
      },
      {
        heading: '下一步',
        body: '把高频排障步骤沉淀成 playbook，让值班同学可以快速执行。',
      },
    ],
  },
]

export const projectCases: ProjectCase[] = [
  {
    id: 'edge-observability',
    name: 'Edge Observability Dashboard',
    summary: '面向边缘节点的统一可观测看板，聚焦延迟与稳定性治理。',
    tech: ['Vue', 'Go', 'MySQL', 'Redis'],
    role: '全栈开发 / 架构设计',
    status: 'Active',
    challenge: '节点分布广、网络波动大，传统中心化监控难以及时发现局部异常。',
    solution: '设计分层采集与聚合模型，将节点健康信号统一到单一视图，并配置告警分级策略。',
    impact: '故障定位平均时间缩短，值班响应效率提升，运维透明度显著改善。',
    links: [
      { label: 'GitHub', href: 'https://github.com/' },
      { label: 'Case Study', href: '/projects/edge-observability' },
    ],
  },
  {
    id: 'knowledge-hub',
    name: 'Engineering Knowledge Hub',
    summary: '把规范、架构和变更决策沉淀为可检索的工程知识库。',
    tech: ['Markdown', 'Vue', 'Node.js'],
    role: '产品定义 / 前端开发',
    status: 'Maintained',
    challenge: '团队知识分散在聊天记录和口头同步中，重复讨论成本高。',
    solution: '按规则、规格、决策、复盘建立目录结构，并将更新流程嵌入日常开发动作。',
    impact: '新人上手时间缩短，关键背景可追溯，跨人协作更稳定。',
    links: [
      { label: 'GitHub', href: 'https://github.com/' },
      { label: 'Case Study', href: '/projects/knowledge-hub' },
    ],
  },
  {
    id: 'blog-experience-kit',
    name: 'Blog Reading Experience Kit',
    summary: '面向技术写作的阅读体验组件集，强调信息层次和可读性。',
    tech: ['Vue', 'CSS Tokens'],
    role: 'UI 设计 / 前端实现',
    status: 'Draft',
    challenge: '文章页易出现层级混乱、正文密度失衡，影响阅读完成率。',
    solution: '统一标题体系、正文容器、目录导航和文末推荐模块，实现一致的阅读流。',
    impact: '文章结构更清晰，读者停留时长提升，后续扩展专栏更顺畅。',
    links: [
      { label: 'GitHub', href: 'https://github.com/' },
      { label: 'Case Study', href: '/projects/blog-experience-kit' },
    ],
  },
]

export const getBlogPostBySlug = (slug: string): BlogPost | undefined => {
  return blogPosts.find((post) => post.slug === slug)
}

export const getProjectById = (id: string): ProjectCase | undefined => {
  return projectCases.find((item) => item.id === id)
}
