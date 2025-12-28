exception NoArgs

fun printWithNumbers(lines, counter) =
    case (lines) of
    (x::[]) => 
        if x = ""
            then print("")
            else print(String.concat([Int.toString(counter), " ", x]))
    |(x::y::xs) =>
        if String.isSuffix "\n" x
        then (print(String.concat([Int.toString(counter), " ", x])); printWithNumbers(y::xs, counter + 1))
        else (print(String.concat([Int.toString(counter), " ", x])); print(y); printWithNumbers(xs, counter + 1))
    |_ => print("")

fun printWithNumbersNotEmptyLines(lines, counter) =
    case (lines) of
    ("\n"::xs) => (print("\n"); printWithNumbersNotEmptyLines(xs, counter))
    |(x::[]) => 
        if x = ""
            then print("")
            else print(String.concat([Int.toString(counter), " ", x]))
    |(x::y::xs) =>
        if String.isSuffix "\n" x
        then (print(String.concat([Int.toString(counter), " ", x])); printWithNumbersNotEmptyLines(y::xs, counter + 1))
        else (print(String.concat([Int.toString(counter), " ", x])); print(y); printWithNumbersNotEmptyLines(xs, counter + 1))
    |_ => print("")

fun collectLines(streams) =
    let 
        fun collectALine(stream, currentLine, lines) =
            case (TextIO.endOfStream(stream), TextIO.input1(stream), TextIO.lookahead(stream), currentLine, lines) of
                (true, _, _, cl, ls) => (TextIO.closeIn(stream); cl::ls)
                |(false, SOME c, _, cl, ls) =>
                    if String.str(c) = "\n"  
                        then collectALine(stream, "", (cl ^ String.str(c))::ls)
                        else collectALine(stream, (cl ^ String.str(c)), ls) 
                |(_,_,_, cl, ls) => collectALine(stream, cl, ls)
    in
        List.rev (List.concat (List.rev(List.map (fn (x) => collectALine(x, "", [])) streams)))
    end

fun cat() = 
    let
    val arguments =  CommandLine.arguments ()
        val (streams, withNumbers, withNumbersNoEmptyLines) = case arguments of
            [] => ([TextIO.stdIn], false, false)
            |"-n"::[] => ([TextIO.stdIn], true, false)
            |"-nb"::[] => ([TextIO.stdIn], true, true)
            |"-bn"::[] => ([TextIO.stdIn], true, true)
            |"-b"::[] => ([TextIO.stdIn], true, true)
            |"-"::[] => ([TextIO.stdIn], false, false)
            |x::[] => ([TextIO.openIn x], false, false)
            |"-n"::xs => ((List.map (fn (r) => if r = "-" then TextIO.stdIn else TextIO.openIn r) xs), true, false)
            |"-bn"::xs => ((List.map (fn (r) => if r = "-" then TextIO.stdIn else TextIO.openIn r) xs), true, true)
            |"-nb"::xs => ((List.map (fn (r) => if r = "-" then TextIO.stdIn else TextIO.openIn r) xs), true, true)
            |"-b"::xs => ((List.map (fn (r) => if r = "-" then TextIO.stdIn else TextIO.openIn r) xs), true, true)
            |x => ((List.map (fn (r) => if r = "-" then TextIO.stdIn else TextIO.openIn r) x), false, false)
    in
        if withNumbers andalso withNumbersNoEmptyLines
            then printWithNumbersNotEmptyLines((collectLines(streams)), 1)
            else if withNumbers
            then printWithNumbers((collectLines(streams)), 1)
            else List.app (fn (x) => print(x)) (collectLines streams)
    end

fun main () = cat()
val _ = main ();
