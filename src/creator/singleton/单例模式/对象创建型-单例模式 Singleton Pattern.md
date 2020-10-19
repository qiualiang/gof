# 对象创建型-单例模式 Singleton Pattern

在有些系统中，为了节省内存资源、保证数据内容的一致性，对某些类要求只能创建一个实例，这就是所谓的单例模式。

单例模式的定义与特点
----------

单例（Singleton）模式的定义：指一个类只有一个实例，且该类能自行创建这个实例的一种模式。例如，Windows 中只能打开一个任务管理器，这样可以避免因打开多个任务管理器窗口而造成内存资源的浪费，或出现各个窗口显示内容的不一致等错误。

在计算机系统中，还有 Windows 的回收站、操作系统中的文件系统、多线程中的线程池、显卡的驱动程序对象、打印机的后台处理服务、应用程序的日志对象、数据库的连接池、网站的计数器、Web 应用的配置对象、应用程序中的对话框、系统中的缓存等常常被设计成单例。

单例模式在现实生活中的应用也非常广泛，例如公司 CEO、部门经理等都属于单例模型。J2EE 标准中的 [Servlet](http://c.biancheng.net/servlet/)Context 和 ServletContextConfig、[Spring](http://c.biancheng.net/spring/) 框架应用中的 ApplicationContext、数据库中的连接池等也都是单例模式。

单例模式有 3 个特点：

1. 单例类只有一个实例对象；
2. 该单例对象必须由单例类自行创建；
3. 单例类对外提供一个访问该单例的全局访问点。

单例模式的优点和缺点
----------

单例模式的优点：

* 单例模式可以保证内存里只有一个实例，减少了内存的开销。
* 可以避免对资源的多重占用。
* 单例模式设置全局访问点，可以优化和共享资源的访问。

单例模式的缺点：

* 单例模式一般没有接口，扩展困难。如果要扩展，则除了修改原来的代码，没有第二种途径，违背开闭原则。
* 在并发测试中，单例模式不利于代码调试。在调试过程中，如果单例中的代码没有执行完，也不能模拟生成一个新的对象。
* 单例模式的功能代码通常写在一个类中，如果功能设计不合理，则很容易违背单一职责原则。

> 单例模式看起来非常简单，实现起来也非常简单。单例模式在面试中是一个高频面试题。希望大家能够认真学习，掌握单例模式，提升核心竞争力，给面试加分，顺利拿到 Offer。

单例模式的应用场景
---------

对于 [Java](http://c.biancheng.net/java/) 来说，单例模式可以保证在一个 JVM 中只存在单一实例。单例模式的应用场景主要有以下几个方面。

* 需要频繁创建的一些类，使用单例可以降低系统的内存压力，减少 GC。
* 某类只要求生成一个对象的时候，如一个班中的班长、每个人的身份证号等。
* 某些类创建实例时占用资源较多，或实例化耗时较长，且经常使用。
* 某类需要频繁实例化，而创建的对象又频繁被销毁的时候，如多线程的线程池、网络连接池等。
* 频繁访问数据库或文件的对象。
* 对于一些控制硬件级别的操作，或者从系统上来讲应当是单一控制逻辑的操作，如果有多个实例，则系统会完全乱套。
* 当对象需要被共享的场合。由于单例模式只允许创建一个对象，共享该对象可以节省内存，并加快对象访问速度。如 Web 中的配置对象、数据库的连接池等。

单例模式的结构与实现
----------

单例模式是[设计模式](http://c.biancheng.net/design_pattern/)中最简单的模式之一。通常，普通类的构造函数是公有的，外部类可以通过“new 构造函数()”来生成多个实例。但是，如果将类的构造函数设为私有的，外部类就无法调用该构造函数，也就无法生成多个实例。这时该类自身必须定义一个静态私有实例，并向外提供一个静态的公有函数用于创建或获取该静态私有实例。

下面来分析其基本结构和实现方法。

### 1\. 单例模式的结构

单例模式的主要角色如下。

* 单例类：包含一个实例且能自行创建这个实例的类。
* 访问类：使用单例的类。

其结构如图 1 所示。

![](resources/A66710A1505E82D78FF8EFF6BD2FA9A2.jpg)

### 2\. 单例模式的实现

Singleton 模式通常有两种实现形式。

#### 第 1 种：懒汉式单例

该模式的特点是类加载时没有生成单例，只有当第一次调用 getlnstance 方法时才去创建这个单例。代码如下：

```sh
public class LazySingleton
{
    private static volatile LazySingleton instance=null;    //保证 instance 在所有线程中同步
    private LazySingleton(){}    //private 避免类在外部被实例化
    public static synchronized LazySingleton getInstance()
    {
        //getInstance 方法前加同步
        if(instance==null)
        {
            instance=new LazySingleton();
        }
        return instance;
    }
}
```

注意：如果编写的是多线程程序，则不要删除上例代码中的关键字 volatile 和 synchronized，否则将存在线程非安全的问题。如果不删除这两个关键字就能保证线程安全，但是每次访问时都要同步，会影响性能，且消耗更多的资源，这是懒汉式单例的缺点。

go 语言版 懒汉式实现

```sh
package main

import (
	"fmt"
	"sync"
)

type HungrySingleton struct {
	name string
}

var instance *HungrySingleton

func (singleton *HungrySingleton) String() string {
	return singleton.name
}
func (singleton *HungrySingleton) GetName() string {
	return singleton.name
}
func (singleton *HungrySingleton) SetName(name string) {
	singleton.name = name
}

func (singleton *HungrySingleton) GetInstance() *HungrySingleton {
	var mu sync.Mutex
	mu.Lock()
	if instance == nil {
		fmt.Println("Creating instance!")
		instance = &HungrySingleton{name: "Gina"}
	} else {
		fmt.Println("Instance has been created!")
	}
	mu.Unlock()
	return instance
}
func main() {
	fmt.Printf("original instance:%s\n", instance)
	instance1 := new(HungrySingleton).GetInstance()
	instance2 := new(HungrySingleton).GetInstance()
	if instance1 == instance2 {
		fmt.Println("instance1 is equal to instance2")
	} else {
		fmt.Println("instance1 isn't equal to instance2")
	}
	fmt.Printf("instance1:%s,instance2:%s\n", instance1, instance2)
}

```

#### 第 2 种：饿汉式单例

该模式的特点是类一旦加载就创建一个单例，保证在调用 getInstance 方法之前单例已经存在了。

```sh
public class HungrySingleton
{
    private static final HungrySingleton instance=new HungrySingleton();
    private HungrySingleton(){}
    public static HungrySingleton getInstance()
    {
        return instance;
    }
}
```

Go 语言实现饿汉模式，代码如下：

```sh
package main

import (
	"fmt"
)

type HungrySingleton struct {
	name string
}

var instance *HungrySingleton

func (singleton *HungrySingleton) String() string {
	return singleton.name
}
func (singleton *HungrySingleton) GetName() string {
	return singleton.name
}
func (singleton *HungrySingleton) SetName(name string) {
	singleton.name = name
}

func (singleton *HungrySingleton) GetInstance() *HungrySingleton {
	return instance
}
func init() {
	instance = &HungrySingleton{name: "gina"}
}
func main() {
	fmt.Printf("original instance:%s\n", instance)
	instance1 := new(HungrySingleton).GetInstance()
	instance2 := new(HungrySingleton).GetInstance()

	if instance1 == instance2 {
		fmt.Println("instance1 is equal to instance2")
	} else {
		fmt.Println("instance1 isn't equal to instance2")
	}
	instance1.SetName("haha")
	fmt.Printf("after change name :instance1:%s,instance2:%s\n", instance1, instance2)
}


```

饿汉式单例在类创建的同时就已经创建好一个静态的对象供系统使用，以后不再改变，所以是线程安全的，可以直接用于多线程而不会出现问题。

单例模式的应用实例
---------

【例1】用懒汉式单例模式模拟产生美国当今总统对象。

分析：在每一届任期内，美国的总统只有一人，所以本实例适合用单例模式实现，图 2 所示是用懒汉式单例实现的结构图。

![](resources/53EEC13443B1CF7751D3D763D35AF6CA.jpg)

```sh
public class SingletonLazy
{
    public static void main(String[] args)
    {
        President zt1=President.getInstance();
        zt1.getName();    //输出总统的名字
        President zt2=President.getInstance();
        zt2.getName();    //输出总统的名字
        if(zt1==zt2)
        {
           System.out.println("他们是同一人！");
        }
        else
        {
           System.out.println("他们不是同一人！");
        }
    }
}
class President
{
    private static volatile President instance=null;    //保证instance在所有线程中同步
    //private避免类在外部被实例化
    private President()
    {
        System.out.println("产生一个总统！");
    }
    public static synchronized President getInstance()
    {
        //在getInstance方法上加同步
        if(instance==null)
        {
               instance=new President();
        }
        else
        {
           System.out.println("已经有一个总统，不能产生新总统！");
        }
        return instance;
    }
    public void getName()
    {
        System.out.println("我是美国总统：特朗普。");
    }  
}
```

程序运行结果如下：

    产生一个总统！
    我是美国总统：特朗普。
    已经有一个总统，不能产生新总统！
    我是美国总统：特朗普。
    他们是同一人

单例模式的扩展
-------

单例模式可扩展为有限的多例（Multitcm）模式，这种模式可生成有限个实例并保存在 ArrayList 中，客户需要时可随机获取，其结构图如图 5 所示。

![](resources/7CDA8FE400B1EEE2BFA626C2DDC42746.jpg)



几种实现方式：

懒汉式，线程不安全，延迟初始化

![](resources/D43276D44BE5B3F2ADA5CCE5CF36BA59.jpg)

懒汉式，线程安全，延迟初始化

![](resources/85D2CCC6F0EDE83484647BEFA2CB42A1.jpg)

饿汉式，线程安全，非延迟初始化，比较常用

![](resources/FB4A7A0B7239828319FA9CE480DE1615.jpg)

双检锁DCL，线程安全，延迟初始化

![](resources/7772F94C1F234985D8165127614E76F0.jpg)

登记式，静态内部类，延迟初始化，线程安全

![](resources/19F1021F68D0D2F527D295B40EA744B6.jpg)

![单例模式的 UML 图](resources/EDD5A30832D195AAB43B0F0C9FDBCFE1.jpg)

