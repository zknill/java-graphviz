# java graphviz

Parse java code into a graphviz image.

The cli tool takes a directory and recurses to find all the java file. It then builds a graphviz output of those files. 

## Usage

Generate a graphviz graph from java files
```bash
java-graphviz -dir /directory/of/source/java/files 
```

Filter the output to a package prefix
```bash
java-graphviz -dir /directory/of/source/java/files -filter "com.site.my.package"
```

Color the output graph by node and it's links
```bash
java-graphviz -dir /directory/of/source/java/files -filter "com.site.my.package" -colorMap "com.site.my.package.MyClass=red"
```

## Example

This is an example from the `parser` test data directory:

```bash
java-graphviz -dir parser/fixtures/ -colorMap="parser.fixtures.single.MyClass=red"
```

![output](https://github.com/zknill/java-graphviz/blob/main/parser-test-output.png)
