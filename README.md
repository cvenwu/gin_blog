# gin_blog
> gin + vue 全栈制作一个博客



## 使用到的第三方包：
1. [gin](https://gin-gonic.com/zh-cn/docs/quickstart/)
2. [gorm](https://gorm.io/zh_CN/)
3. [ini](https://ini.unknwon.io/docs/intro/getting_started)
4. 加密使用：scrypt(几乎不可以破解，区块链中用的多，**但是开销不知道，如果要正式的话一定要稳重，用成熟的方案比较好**)或bcrypt


## 去做的
1. 使用gorm的钩子函数，在保存一个user之前，执行一个函数，应该是beforesave吧
