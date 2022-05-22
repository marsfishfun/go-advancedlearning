# go-advancedlearning
Go进阶训练营的学习

# 作业提交URL
Week02：https://jinshuju.net/f/FMuVMP
Week03：https://jinshuju.net/f/qQt8Gd
Week04：https://jinshuju.net/f/DL2cP7
Week05：https://jinshuju.net/f/FL8nyE
Week08：https://jinshuju.net/f/ugThQB
Week09：https://jinshuju.net/f/A8G5Lv
毕业项目&总结：https://jinshuju.net/f/lMfuLL

# go-advancedlearning
Go进阶训练营的学习

# Week02
dao层中遇到的错误，是否应该Wrap这个error，抛给上层，为什么，应该怎么做？
应该往上抛，内部无法降级处理的error，抛给上层处理
