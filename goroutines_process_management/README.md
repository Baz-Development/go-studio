## Context

Softwares geralmente iniciam processos de longa duração e de uso intensivo de recursos (muitas vezes em goroutines). Se a ação que causou isso é cancelada ou falha por algum motivo, você precisa parar esses processos de forma consistente dentro da sua aplicação.

Se você não gerenciar isso, sua aplicação Go tão ágil da qual você tem tanto orgulho pode começar a ter problemas de desempenho difíceis de depurar.