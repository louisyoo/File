# File
提供用于创建、复制、删除、移动和打开单一文件的静态方法

##函数列表

1. Exists(filename string)(bool)
确定指定的文件是否存在。
2. ReadAllText(filename string)(string,err)
打开一个文本文件，读取文件的所有字符串，然后关闭该文件
3. ReadAllBytes(filename string)([]byte,err)
打开一个文本文件，读取文件的所有字符串，然后关闭该文件
